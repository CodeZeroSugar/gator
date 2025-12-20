package main

import (
	"context"
	"fmt"
)

func reset(s *state, cmd command) error {
	fmt.Println("running reset...")
	if len(cmd.args) > 0 {
		return fmt.Errorf("too many args for reset")
	}

	tx, err := s.dbPool.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	qtx := s.db.WithTx(tx)

	/*if err := qtx.ResetFeedFollows(context.Background()); err != nil {
		return fmt.Errorf("couldn't reset feed follows in transaction: %w", err)
	}
	*/
	if err := qtx.ResetFeeds(context.Background()); err != nil {
		return fmt.Errorf("couldn't reset feeds in transaction: %w", err)
	}

	if err := qtx.ResetUsers(context.Background()); err != nil {
		return fmt.Errorf("couldn't reset users in transaction: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("database reset was sucessful")
	return nil
}
