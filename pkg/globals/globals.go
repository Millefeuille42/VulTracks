package globals

var Dev = true

var (
	RootLocation     = getRootLocation()
	TemplateLocation = RootLocation + "/templates"
	StaticLocation   = RootLocation + "/static"
	ConfigLocation   = RootLocation + "/config"

	DatabaseLocation = RootLocation + "/database.sqlite3"

	FirstRun = false
	Shutdown = false
)
