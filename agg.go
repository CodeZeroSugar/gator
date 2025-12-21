package main

import (
	"fmt"
	"time"
)

func agg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("agg takes one argument")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to parse time between requests: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s.db); err != nil {
			return fmt.Errorf("error returned from scrape feeds: %w", err)
		}
	}
}
