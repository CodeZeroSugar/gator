package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/CodeZeroSugar/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("error: name must be entered")
	}
	inputName := cmd.args[0]
	newUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      inputName,
	})
	if err != nil {
		log.Fatal("error: user already exists, exiting...")
	}

	if err := s.configPtr.SetUser(inputName); err != nil {
		return fmt.Errorf("error: failed to set user after creation")
	}
	fmt.Printf("new user '%v' was created\n", inputName)
	fmt.Printf("%+v", newUser)

	return nil
}
