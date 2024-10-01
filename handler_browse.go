package main

import (
	"fmt"
	"errors"
	"context"
	"strconv"
	"github.com/P3T3R2002/blog_aggreGATOR/internal/database"
)

func handlerBrowse(s *State, cmd Command, user database.User) error {
	var limit int = 2
	if len(cmd.args) > 2 {
		return errors.New(fmt.Sprintf("Need zero or one argument for %s: cli <command> [args...]", cmd.name))
	} else if len(cmd.args) == 1 {
		l, err := strconv.Atoi(cmd.args[0])
		if err !=  nil {
			return err
		}
		limit = l
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{user.ID, limit})
	if err != nil {
		return err
	}
	printPosts(posts)
	return nil
}

func printPosts(posts []database.Post) {
	for _, post := range posts {
		fmt.Println("//----------------------")
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Published at: %v\n", post.PublishedAt)
		fmt.Printf("Created at: %v\n", post.CreatedAt)
		fmt.Printf("Updated at: %v\n", post.UpdatedAt)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Printf("Description: %s\n", post.Description)
	}
}