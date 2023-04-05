package models

import (
	"VulTracks/pkg/database"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type TrackModel struct {
	Id     string `json:"id"`
	Path   string `json:"path" validate:"required"`
	Name   string `json:"name" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}

func (track *TrackModel) CreateTable() error {
	_, err := database.Database.Exec(`
		CREATE TABLE IF NOT EXISTS tracks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			path TEXT NOT NULL UNIQUE,
			user_id INTEGER NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id),
			name TEXT NOT NULL
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func appendTracksToList(list []TrackModel, rows *sql.Rows) ([]TrackModel, error) {
	var track TrackModel
	err := rows.Scan(&track.Id, &track.Path, &track.UserId, &track.Name)
	if err != nil {
		return nil, err
	}
	list = append(list, track)
	return list, nil
}

func getTracksListFromRows(rows *sql.Rows) ([]TrackModel, error) {
	var list []TrackModel
	var err error
	list, err = appendTracksToList(list, rows)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		list, err = appendTracksToList(list, rows)
		if err != nil {
			return nil, err
		}
	}
	return list, nil
}

func (track *TrackModel) getTrackByQuery(query squirrel.SelectBuilder) error {
	rows, err := database.SelectHelper(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	err = rows.Scan(&track.Id, &track.Path, &track.UserId, &track.Name)
	if err != nil {
		return err
	}
	return nil
}

func GetTracks() ([]TrackModel, error) {
	rows, err := database.SelectHelper(
		squirrel.
			Select("*").
			From("tracks"),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tracks, err := getTracksListFromRows(rows)
	return tracks, err
}

func (track *TrackModel) GetTrackById(id string) (*TrackModel, error) {
	return track, track.getTrackByQuery(
		squirrel.
			Select("*").
			From("tracks").
			Where(squirrel.Eq{"id": id}),
	)
}

func GetTracksByUserId(userId string) ([]TrackModel, error) {
	rows, err := database.SelectHelper(
		squirrel.
			Select("*").
			From("tracks").
			Where(squirrel.Eq{"user_id": userId}),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return getTracksListFromRows(rows)
}

func (track *TrackModel) CreateTrack() error {
	result, err := squirrel.
		Insert("tracks").
		Columns("path", "user_id", "name").
		Values(track.Path, track.UserId, track.Name).
		RunWith(database.Database).Exec()
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	track.Id = fmt.Sprintf("%d", id)
	return nil
}

func (track *TrackModel) DeleteTrack() error {
	_, err := squirrel.
		Delete("tracks").
		Where(squirrel.Eq{"id": track.Id}).
		RunWith(database.Database).Exec()
	return err
}

func (track *TrackModel) UpdateTrack() error {
	trackQuery := squirrel.Update("tracks")

	if track.Path != "" {
		trackQuery = trackQuery.Set("path", track.Path)
	}

	if track.UserId != "" {
		trackQuery = trackQuery.Set("user_id", track.UserId)
	}

	if track.Name != "" {
		trackQuery = trackQuery.Set("name", track.Name)
	}

	_, err := trackQuery.
		Where(squirrel.Eq{"id": track.Id}).
		RunWith(database.Database).Exec()
	return err
}

func (track *TrackModel) GetNumberOfTrackByUserId(userId string) (int, error) {
	rows, err := database.SelectHelper(
		squirrel.Select("COUNT(*)").
			From("tracks").
			Where(squirrel.Eq{"user_id": userId}),
	)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
