package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/config"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/handlers"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/middleware"
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
	cmds.Register("login", handlers.HandlerLogin)
	cmds.Register("register", handlers.HandlerRegister)
	cmds.Register("reset", handlers.HandlerReset)
	cmds.Register("users", handlers.HandlerUsers)
	cmds.Register("agg", handlers.HandlerAgg)
	cmds.Register("addfeed", middleware.LoggedIn(handlers.HandlerAddFeed))
	cmds.Register("feeds", handlers.HandlerFeeds)
	cmds.Register("follow", middleware.LoggedIn(handlers.HandlerFollow))
	cmds.Register("following", middleware.LoggedIn(handlers.HandlerFollowing))
	cmds.Register("unfollow", middleware.LoggedIn(handlers.HandlerUnfollow))
	cmds.Register("browse", middleware.LoggedIn(handlers.HandlerBrowse))
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
