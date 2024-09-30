package main

import (
	"fmt"
	"errors"
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/P3T3R2002/blog_aggreGATOR/internal/database"
)

func handlerFollow(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return errors.New(fmt.Sprintf("Need only one argument for %s: cli <command> [args...]", cmd.name))
	}
	url := cmd.args[0]

	
	feed_ID, err := s.db.GetFeedID(context.Background(), url)
	if err != nil {
		return err
	}
	
	user_ID, err := s.db.GetUserID(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return err
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), 
		database.CreateFeedFollowParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user_ID,
			FeedID:    feed_ID,
		})
	if err != nil {
		return err
	}

	fmt.Println(follow.UserName, follow.FeedName)

	fmt.Println("Followed successfully!")
	return nil
}

func handlerFollowing(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return errors.New(fmt.Sprintf("Need zero argument for %s: cli <command> [args...]", cmd.name))
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}
	return nil
}