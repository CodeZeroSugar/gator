package main

import (
	"context"
	"fmt"
)

func following(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("following command does not take arguments")
	}
	currentUser, err := s.db.GetUser(context.Background(), s.configPtr.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could not get current user from database: %w", err)
	}

	followingFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("could not get feed follows: %w", err)
	}

	fmt.Printf("Feeds followed by %v:\n", currentUser.Name)
	for _, f := range followingFeeds {
		feedRow := f.FeedName
		fmt.Printf("- %v\n", feedRow)
	}

	return nil
}
