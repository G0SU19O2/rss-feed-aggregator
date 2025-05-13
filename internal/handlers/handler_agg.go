package handlers

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/rss"
)

func HandlerAgg(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("only support one argument")
	}
	time_between_reqs := cmd.Args[0]
	timeBetweenRequests, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("can't parse time")
	}
	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests.String())
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		error := rss.ScrapeFeeds(context.Background(), s.Db)
		if error != nil {
			fmt.Println(error.Error())
			os.Exit(1)
		}
	}
}
