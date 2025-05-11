package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/config"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, _ := config.Read()
	db, err := sql.Open("mysql", cfg.Db_url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	programState := &cli.State{
		Cfg: &cfg,
		Db:  dbQueries,
	}

	cmds := cli.Commands{
		RegisteredCommands: make(map[string]func(*cli.State, cli.Command) error),
	}
	cmds.Register("login", handlerLogin)
	cmds.Register("register", handlerRegister)
	cmds.Register("reset", handlerReset)
	cmds.Register("users", handlerUsers)
	cmds.Register("agg", handlerAgg)
	cmds.Register("addfeed", handlerAddFeed)
	cmds.Register("feeds", handlerFeeds)
	cmds.Register("follow", handlerFollow)
	cmds.Register("following", handlerFollowing)
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	err = cmds.Run(programState, cli.Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
