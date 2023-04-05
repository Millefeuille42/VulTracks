package models

import (
	"VulTracks/pkg/database"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Id       string `json:"id"`
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8"`
}

func (user *UserModel) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *UserModel) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user *UserModel) CreateTable() error {
	_, err := database.Database.Exec(`
		CREATE TABLE IF NOT EXISTS users (
		    			id INTEGER PRIMARY KEY AUTOINCREMENT,
		    			username TEXT NOT NULL UNIQUE,
		    			password TEXT NOT NULL
		    		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func appendUserToList(list []UserModel, rows *sql.Rows) ([]UserModel, error) {
	var user UserModel
	err := rows.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	list = append(list, user)
	return list, nil
}

func getUserListFromRows(rows *sql.Rows) ([]UserModel, error) {
	var list []UserModel
	var err error
	list, err = appendUserToList(list, rows)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		list, err = appendUserToList(list, rows)
		if err != nil {
			return nil, err
		}
	}
	return list, nil
}

func (user *UserModel) getUserByQuery(query squirrel.SelectBuilder) error {
	rows, err := database.SelectHelper(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	err = rows.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]UserModel, error) {
	rows, err := database.SelectHelper(
		squirrel.
			Select("*").
			From("users"),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return getUserListFromRows(rows)
}

func (user *UserModel) GetUserByUsername(username string) (*UserModel, error) {
	return user, user.getUserByQuery(
		squirrel.
			Select("*").
			From("users").
			Where(squirrel.Eq{"username": username}).
			Limit(1),
	)
}

func (user *UserModel) GetUserById(id string) (*UserModel, error) {
	return user, user.getUserByQuery(
		squirrel.
			Select("*").
			From("users").
			Where(squirrel.Eq{"id": id}).
			Limit(1),
	)
}

func (user *UserModel) GetUserByIdOrUsername(idOrUsername string) (*UserModel, error) {
	return user, user.getUserByQuery(
		squirrel.
			Select("*").
			From("users").
			Where(squirrel.Or{squirrel.Eq{"id": idOrUsername}, squirrel.Eq{"username": idOrUsername}}).
			Limit(1),
	)
}

func (user *UserModel) CreateUser() error {
	result, err := squirrel.
		Insert("users").
		Columns("username", "password").
		Values(user.Username, user.Password).
		RunWith(database.Database).Exec()
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = fmt.Sprintf("%d", id)
	return nil
}

func (user *UserModel) DeleteUser() error {
	_, err := squirrel.
		Delete("users").
		Where(squirrel.Eq{"id": user.Id}).
		RunWith(database.Database).Exec()
	return err
}

func (user *UserModel) UpdateUser(newUser *UserModel) error {
	userQuery := squirrel.Update("users")

	if newUser.Password != "" {
		err := newUser.HashPassword()
		if err != nil {
			return err
		}
		userQuery.Set("password", newUser.Password)
		user.Password = newUser.Password
	}

	if newUser.Username != "" {
		userQuery.Set("username", newUser.Username)
		user.Username = newUser.Username
	}

	_, err := userQuery.
		Where(squirrel.Eq{"id": user.Id}).
		RunWith(database.Database).Exec()
	return err
}
