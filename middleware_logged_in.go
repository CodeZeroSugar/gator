package main

import (
	"context"
	"fmt"

	"github.com/CodeZeroSugar/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.configPtr.CurrentUserName)
		if err != nil {
			return fmt.Errorf("inner func of middlewareLoggedIn failed to get current user: %w", err)
		}
		return handler(s, cmd, currentUser)
	}
}
