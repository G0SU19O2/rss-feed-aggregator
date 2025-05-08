package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/config"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
	_ "github.com/go-sql-driver/mysql"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, _ := config.Read()
	db, err := sql.Open("mysql", cfg.Db_url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	programState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addFeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
