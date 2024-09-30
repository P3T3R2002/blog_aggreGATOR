package main

import (
	"fmt"
	"errors"
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/P3T3R2002/blog_aggreGATOR/internal/database"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return errors.New(fmt.Sprintf("Need only one argument for %s: cli <command> [args...]", cmd.name))
	}
	name := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return errors.New(fmt.Sprintf("User not registered: %s", name))
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't set current user: %w", err))
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return errors.New(fmt.Sprintf("Need only one argument for %s: cli <command> [args...]", cmd.name))
	}
	name := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return errors.New(fmt.Sprintf("User already registered: %s", name))
	}

	_, err = s.db.CreateUser(context.Background(), 
		database.CreateUserParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: name,
		})
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't set current user: %w", err))
	}

	fmt.Println("User created successfully!")
	return nil
}

func handlerReset(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return errors.New(fmt.Sprintf("Need zero argument for %s: cli <command> [args...]", cmd.name))
	}

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Database reset successfully!")
	return nil
}

func handlerUsers(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return errors.New(fmt.Sprintf("Need zero argument for %s: cli <command> [args...]", cmd.name))
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user == s.cfg.Current_user_name {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}
	}
	return nil
}
