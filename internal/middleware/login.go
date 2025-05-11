package middleware

import (
	"context"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
)

func LoggedIn(handler func(s *cli.State, cmd cli.Command, user database.User) error) func(*cli.State, cli.Command) error {
	return func(s *cli.State, cmd cli.Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
