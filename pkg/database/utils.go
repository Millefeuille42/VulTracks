package database

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func SelectHelper(query squirrel.SelectBuilder) (*sql.Rows, error) {
	rows, err := query.RunWith(Database).Query()
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		rows.Close()
		return nil, fiber.ErrNotFound
	}

	return rows, nil
}
