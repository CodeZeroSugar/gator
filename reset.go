package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

func reset(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		log.Fatal("too many args for reset, exiting...")
	}
	if err := s.db.Reset(context.Background()); err != nil {
		log.Fatal("reset failed, exiting...")
	}
	fmt.Println("database reset was sucessful")
	os.Exit(0)
	return nil
}
