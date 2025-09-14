package data

import (
    "database/sql"
    "errors"
)

// custom error - returned when no movie is found in DB.
var (
	ErrRecordNotFound = errors.New("record not found.")
)

type Models struct {
    Movies MovieModel
}

func NewModels (db *sql.DB) Models {
	return Models {
		Movies: MovieModel{DB: db},
	}
}