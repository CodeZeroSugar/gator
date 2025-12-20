package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/CodeZeroSugar/gator/internal/database"
	"github.com/google/uuid"
)

func addFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		log.Fatal("not enough args, need 2...")
	}
	currUser := s.configPtr.CurrentUserName
	dbUser, err := s.db.GetUser(context.Background(), currUser)
	if err != nil {
		log.Fatal(err)
	}
	userID := dbUser.ID
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    userID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("add feed failed to create feed: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("add feed failed to create feed follow: %w", err)
	}

	fmt.Printf("%+v", feed)

	return nil
}
