package main

import (
	"github.com/CodeZeroSugar/gator/internal/config"
	"github.com/CodeZeroSugar/gator/internal/database"
)

type state struct {
	db        *database.Queries
	configPtr *config.Config
}
