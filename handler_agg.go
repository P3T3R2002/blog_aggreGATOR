package main

import (
	"fmt"
	"errors"
	"time"
	"context"
	"log"
	"database/sql"
	"github.com/google/uuid"
	"github.com/P3T3R2002/blog_aggreGATOR/internal/database"
)

func handlerAgg(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return errors.New(fmt.Sprintf("Need 1 argument for %s: cli <command> [args...]", cmd.name))
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid time string: %w\n", err)) 
	}
	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	return nil
}

func scrapeFeeds(s *State) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("Couldn't fetch next feed: %w\n", err)
		return 
	}
	
	log.Println("Found feed to fetch.")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	feedData, err := fetchFeed(feed.Url)
	if err != nil {
		log.Printf("Couldn't fetch feed %s: %w\n", feed.Name, err)
		return 
	}

	_, err = db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %w\n", feed.Name, err)
		return 
	}

	for _, item := range feedData.Channel.Item {
		title := getNullString(item.Title)
		description := getNullString(item.Description)

		parsedTime, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			log.Println("Error parsing time:", err)
			continue
		}
		
		_, err = db.CreatePost(context.Background(), 
		database.CreatePostParams{
			ID:				uuid.New(),
			CreatedAt:		time.Now(),
			UpdatedAt:		time.Now(),
			Title:			title,
			Url:			item.Link,
			Description:	description,
			PublishedAt:	parsedTime,
			FeedID:			feed.ID,
		})
		if err != nil {
			log.Println("Failed to create post!")
			continue
		}
	}

	log.Printf("%v posts found in feed %s.\n", len(feedData.Channel.Item), feed.Name)
}

func getNullString(s string) sql.NullString {
	var nullString sql.NullString

	if s != "" {
		nullString = sql.NullString{
			String: s,
			Valid:  true,
		}
	} else {
		nullString = sql.NullString{
			Valid: false,
		}
	}
	return nullString
}
