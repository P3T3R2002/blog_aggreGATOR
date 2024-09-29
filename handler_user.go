package main

import (
	"fmt"
	"errors"
)

func handler_Login(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return errors.New(fmt.Sprintf("Need only one argument for %s: cli <command> [args...]", cmd.name))
	}
	name := cmd.args[0]

	err := s.cfg.Set_user(name)
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't set current user: %w", err))
	}

	fmt.Println("User switched successfully!")
	return nil
}
