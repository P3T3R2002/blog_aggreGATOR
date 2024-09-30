package main

import (
	"fmt"
	"errors"
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/P3T3R2002/blog_aggreGATOR/internal/database"
)

func handlerAgg(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return errors.New(fmt.Sprintf("Need zero argument for %s: cli <command> [args...]", cmd.name))
	}

	feed, err := fetchFeed("https://www.wagslane.dev/index.xml")
	if err != nil {
		return errors.New(fmt.Sprintf("couldn't fetch feed: %w\n", err))
	}

	fmt.Printf("Feed: %+v\n", feed)
	return nil
}

func handlerAddFeed(s *State, cmd Command) error {
	if len(cmd.args) != 2 {
		return errors.New(fmt.Sprintf("Need two argument for %s: cli <command> [args...]", cmd.name))
	}
	name := cmd.args[0]
	url := cmd.args[1]
	
	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return errors.New(fmt.Sprintf("Feed already registered: %s", name))
	}

	_, err = s.db.CreateFeed(context.Background(), 
		database.CreateFeedParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: name,
			Url: url,
			UserID: user.ID,
		})
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return err
	}
	
	fmt.Println("Feed created successfully!")

	err = handlerFollow(s, Command{name: "follow", args: cmd.args[1:]})
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}

func handlerFeeds(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return errors.New(fmt.Sprintf("Need zero argument for %s: cli <command> [args...]", cmd.name))
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed)
	}
	return nil
}

