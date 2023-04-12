package main

import (
	"VulTracks/pkg/app"
	"VulTracks/pkg/database"
	"VulTracks/pkg/globals"
	"VulTracks/pkg/middlewares/sessionChecker"
	"VulTracks/pkg/models"
	"VulTracks/pkg/routes/auth"
	"VulTracks/pkg/routes/folder"
	"VulTracks/pkg/routes/index"
	"VulTracks/pkg/routes/settingsRoute"
	"VulTracks/pkg/routes/track"
	"VulTracks/pkg/routes/user"
	"VulTracks/pkg/store"
	"VulTracks/pkg/utils"
	"VulTracks/pkg/utils/id3Utils"
	"VulTracks/pkg/utils/settings"
	"VulTracks/pkg/validator"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"log"
)

func askConsent(sentence string) (bool, error) {
	var consent string
	fmt.Println(sentence + " (y/n)")
	_, err := fmt.Scanln(&consent)
	if err != nil {
		if err.Error() == "unexpected newline" {
			return false, nil
		}
		return false, err
	}
	return consent == "y", nil
}

func initFileSystem() {
	fmt.Println("Installing VulTracks...")

	fmt.Printf("Creating root directory (%s)...\n", globals.RootLocation)
	utils.AutoPanic(utils.CreateDirIfNotExist(globals.RootLocation))

	fmt.Println("Creating database...")
	utils.AutoPanic(utils.CreateFileIfNotExist(globals.DatabaseLocation))
}

func populateDatabase() {
	fmt.Println("Populating database...")
	userModel := models.UserModel{}
	utils.AutoPanic(userModel.CreateTable())
	trackModel := models.TrackModel{}
	utils.AutoPanic(trackModel.CreateTable())
	folderModel := models.FolderModel{}
	utils.AutoPanic(folderModel.CreateTable())
}

func install() {
	consent, err := askConsent("VulTracks is not installed, do you want to install it?")
	utils.AutoPanic(err)
	if !consent {
		return
	}
	initFileSystem()
	database.Database, err = database.NewDatabase("file:" + globals.DatabaseLocation + "?_foreign_keys=ON")
	defer database.Database.Close()
	populateDatabase()
	utils.AutoPanic(settings.ParseDefaults())
	utils.AutoPanic(settings.RewriteSettings())
}

func registerMiddlewares(app *fiber.App) {
	app.Use(logger.New())
	if !globals.Dev {
		app.Use(csrf.New())
	}
	app.Use(favicon.New(favicon.Config{
		File: globals.StaticLocation + "/images/logo.png",
		URL:  "/favicon.ico",
	}))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "VulTracks Metrics Page"}))
	app.Static("/static", globals.StaticLocation)

	app.Use(sessionChecker.New(sessionChecker.Config{
		Filter: sessionChecker.DefaultFilter,
	}))
}

func registerRoutes(app *fiber.App) {
	index.Register(app)
	auth.Register(app)
	user.Register(app)
	track.Register(app)
	folder.Register(app)
	settingsRoute.Register(app)
}

func main() {
	if globals.RootLocation == "" {
		fmt.Println("Unable to find a proper installation location, please set your XDG_CONFIG_HOME environment variable")
		return
	}

	//err := os.RemoveAll(globals.RootLocation)
	//utils.AutoPanic(err)

	var exist bool
	var err error
	if globals.Dev {
		exist, err = utils.IsExist(globals.DatabaseLocation)
	} else {
		exist, err = utils.IsExist(globals.RootLocation)
	}
	utils.AutoPanic(err)
	if !exist {
		install()
		return
	}

	utils.AutoPanic(settings.RefreshSettings())
	utils.AutoPanic(id3Utils.RefreshID3Frames())

	database.Database, err = database.NewDatabase("file:" + globals.DatabaseLocation + "?_foreign_keys=ON")
	defer database.Database.Close()
	utils.AutoPanic(err)
	_, err = models.GetUsers()
	if err != nil {
		if err.Error() == "Not Found" {
			globals.FirstRun = true
			fmt.Println("It seems like this is your first run, go to /onboarding to create your first user")
		} else {
			utils.AutoPanic(err)
		}
	}

	store.Store = store.NewStore()
	validator.Validator = validator.NewValidator()
	app.App = app.NewApp()
	registerMiddlewares(app.App)
	registerRoutes(app.App)
	utils.SetupCloseHandler(app.App)

	for !globals.Shutdown {
		err = app.StartApp(app.App)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
