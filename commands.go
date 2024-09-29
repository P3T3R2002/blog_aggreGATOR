package main

import "errors"

type Command struct {
	name string
	args []string
}

type Commands struct {
	registered_commands map[string]func(*State, Command) error
}

func (c *Commands)register(name string, f func(*State, Command) error) {
	c.registered_commands[name] = f
}

func (c *Commands)run(s *State, cmd Command) error {
	f, ok := c.registered_commands[cmd.name]
	if !ok {
		return errors.New("Command not found!")
	}
	return f(s, cmd)
}

