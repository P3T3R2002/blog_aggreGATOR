package main

import _ "github.com/lib/pq"

import (
	"log"
	"os"
	"database/sql"
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
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFollowing)

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
