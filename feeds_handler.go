package main

import (
	"context"
	"fmt"
)

func feedsHandler(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("the feeds command does not take any arguments")
	}

	feed, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds from database: %w", err)
	}

	fmt.Println("Showing feeds for all users:")
	fmt.Println("")
	for _, row := range feed {
		fName := row.Name
		fURL := row.Url
		uName := row.Name_2.String
		fmt.Printf("Feed Name: %v\n", fName)
		fmt.Printf("Feed URL : %v\n", fURL)
		fmt.Printf("Username : %v\n", uName)
		fmt.Println("")
	}

	return nil
}
