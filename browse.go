package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/CodeZeroSugar/gator/internal/database"
)

func browse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("browse takes at most 1 argument")
	}
	var limit int32
	if len(cmd.args) == 1 {
		limit64, err := strconv.ParseInt(cmd.args[0], 10, 32)
		limit = int32(limit64)
		if err != nil {
			return fmt.Errorf("failed to convert argument to integer: %w", err)
		}
	} else {
		limit = 2
	}

	postsForUserParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), postsForUserParams)
	if err != nil {
		return fmt.Errorf("failed to get posts for user: %w", err)
	}

	fmt.Println("Browsing posts...")
	fmt.Println("")

	for _, post := range posts {
		fmt.Printf("Feed Name: %v\n", post.FeedName)
		fmt.Printf("Title: %v\n", post.Title.String)
		fmt.Printf("Published: %v\n", post.PublishedAt.Time.String())
		fmt.Printf("Description: %v\n", post.Description.String)
		fmt.Println("")

	}
	return nil
}
