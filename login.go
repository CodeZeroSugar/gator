package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("error: username is required")
	}
	userName := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		log.Fatal("error: user does not exists")
	}
	if err := s.configPtr.SetUser(userName); err != nil {
		return fmt.Errorf("error: login handler failed to set the username")
	}
	fmt.Printf("user has been set to %v\n", userName)

	return nil
}
