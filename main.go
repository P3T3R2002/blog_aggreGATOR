package main

import (
	"log"
	"os"
	"github.com/P3T3R2002/blog_aggreGATOR/internal/config"
)

type State struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Reading config error: %v", err)
	}
	
	program_state := &State{
		cfg: &cfg,
	}

	cmds := Commands{
		registered_commands: make(map[string]func(*State, Command) error),
	}
	
	cmds.register("login", handler_Login)

	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments: cli <command> [args...]")
		return
	}

	cmd_name := os.Args[1]
	cmd_args := []string{}
	if len(os.Args) > 2 {
		cmd_args = os.Args[2:]
	}
	err = cmds.run(program_state, Command{name: cmd_name, args: cmd_args})
	if err != nil {
		log.Fatal(err)
	}
} 
