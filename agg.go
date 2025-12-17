package main

import (
	"context"
	"fmt"
	"log"
)

const testURL = "https://www.wagslane.dev/index.xml"

func agg(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		log.Fatal("Too many arguments for agg command")
	}
	feed, err := fetchFeed(context.Background(), testURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", feed)

	return nil
}
