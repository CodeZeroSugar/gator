package main

import (
	"context"
	"fmt"

	"github.com/CodeZeroSugar/gator/internal/database"
)

func unfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("unfollow only takes 1 arg")
	}
	feed, err := s.db.GetFeedIdFromUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("unfollow failed to get feed if from url: %w", err)
	}

	unfollowParams := database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	if err := s.db.UnfollowFeed(context.Background(), unfollowParams); err != nil {
		return fmt.Errorf("failed to unfollow feed: %w", err)
	}

	return nil
}
