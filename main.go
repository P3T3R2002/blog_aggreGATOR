package main

import _ "github.com/lib/pq"

import (
	"log"
	"os"
	"database/sql"
	"context"
	"github.com/P3T3R2002/blog_aggreGATOR/internal/config"
	"github.com/P3T3R2002/blog_aggreGATOR/internal/database"
)

const dbURL = "postgres://postgres:postgres@localhost:5432/gator"

type State struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Opening postgres error: %v", err)
	}
	dbQueries := database.New(db)
	
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Reading config error: %v", err)
	}
	
	programState := &State{
		db: dbQueries,
		cfg: &cfg,
	}

	cmds := Commands{
		registered_commands: make(map[string]func(*State, Command) error),
	}
	
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := []string{}
	if len(os.Args) > 2 {
		cmdArgs = os.Args[2:]
	}
	err = cmds.run(programState, Command{name: cmdName, args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
} 

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}