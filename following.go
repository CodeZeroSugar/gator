package main

import (
	"context"
	"fmt"

	"github.com/CodeZeroSugar/gator/internal/database"
)

func following(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("following command does not take arguments")
	}

	followingFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not get feed follows: %w", err)
	}

	fmt.Printf("Feeds followed by %v:\n", user.Name)
	for _, f := range followingFeeds {
		feedRow := f.FeedName
		fmt.Printf("- %v\n", feedRow)
	}

	return nil
}
