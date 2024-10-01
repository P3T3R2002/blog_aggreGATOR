package main

import (
	"fmt"
	"errors"
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/P3T3R2002/blog_aggreGATOR/internal/database"
)

func handlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.args) != 2 {
		return errors.New(fmt.Sprintf("Need two argument for %s: cli <command> [args...]", cmd.name))
	}
	name := cmd.args[0]
	url := cmd.args[1]

	_, err := s.db.CreateFeed(context.Background(), 
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

	err = handlerFollow(s, Command{name: "follow", args: cmd.args[1:]}, user)
	if err != nil {
		return err
	}

	printFeed(feed)
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

	printFeeds(feeds)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("Feed ID: %v\n", feed.ID)
	fmt.Printf("Created at: %v\n", feed.CreatedAt)
	fmt.Printf("Updated at: %v\n", feed.UpdatedAt)
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("Link: %s\n", feed.Url)
	fmt.Printf("Uploader ID: %v\n", feed.UserID)
	fmt.Printf("Last fetched at: %v\n", feed.LastFetchedAt)
}

func printFeeds(feeds []database.GetFeedsRow) {
	for _, feed := range feeds {
		fmt.Printf("Feed name: %s\n", feed.Name)
		fmt.Printf("Feed link: %s\n", feed.Url)
		fmt.Printf("Created by: %s\n", feed.Name_2)
	}

}