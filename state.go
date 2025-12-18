package main

import (
	"database/sql"

	"github.com/CodeZeroSugar/gator/internal/config"
	"github.com/CodeZeroSugar/gator/internal/database"
)

type state struct {
	db        *database.Queries
	dbPool    *sql.DB
	configPtr *config.Config
}
