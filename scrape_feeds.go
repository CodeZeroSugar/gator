package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/CodeZeroSugar/gator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(d *database.Queries) error {
	nextFeed, err := d.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get next feed: %w", err)
	}
	if _, err := d.MarkFeedFetched(context.Background(), nextFeed.ID); err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %w", err)
	}
	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed during scrape feeds: %w", err)
	}

	for _, item := range feed.Channel.Item {
		title := sql.NullString{
			String: item.Title,
			Valid:  item.Title != "",
		}
		description := sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		}

		var publishedAt sql.NullTime
		if item.PubDate != "" {
			pubTime, err := time.Parse(time.RFC1123Z, item.PubDate)
			if err == nil {
				publishedAt = sql.NullTime{
					Time:  pubTime,
					Valid: true,
				}
			} else {
				fmt.Printf("error parsing published time %q: %v\n", item.PubDate, err)
			}
		}

		feedDB, err := d.GetFeedIdFromUrl(context.Background(), nextFeed.Url)
		if err != nil {
			return fmt.Errorf("could not get feed id from url in scrape feeds: %w", err)
		}
		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       title,
			Url:         item.Link,
			Description: description,
			PublishedAt: publishedAt,
			FeedID:      feedDB.ID,
		}
		_, err1 := d.CreatePost(context.Background(), postParams)
		if err1 != nil {
			continue
		}

	}

	return nil
}
