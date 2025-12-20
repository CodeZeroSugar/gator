package main

import (
	"context"
	"fmt"
	"time"

	"github.com/CodeZeroSugar/gator/internal/database"
	"github.com/google/uuid"
)

func follow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("follow command take 1 arg")
	}
	urlInput := cmd.args[0]
	feedDB, err := s.db.GetFeedIdFromUrl(context.Background(), urlInput)
	if err != nil {
		return fmt.Errorf("could not find feed with provided url: %w", err)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.configPtr.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could not get current user form database: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feedDB.ID,
	}

	feedFollowRecord, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("failed to create feed follow record: %w", err)
	}
	feedName := feedFollowRecord.FeedName
	feedUser := feedFollowRecord.UserName
	fmt.Printf("Feed Name: %v\n", feedName)
	fmt.Printf("Feed User: %v\n", feedUser)

	return nil
}
