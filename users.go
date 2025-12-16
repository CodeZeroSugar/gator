package main

import (
	"context"
	"fmt"
	"log"
)

func users(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		log.Fatal("too many args for users, exiting...")
	}
	dbUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		log.Fatal("error: something went wrong retrieving users, exiting...")
	}
	for _, user := range dbUsers {
		u := user.Name
		if u == s.configPtr.CurrentUserName {
			fmt.Printf("* %v (current)\n", u)
		} else {
			fmt.Printf("* %v\n", u)
		}
	}

	return nil
}
