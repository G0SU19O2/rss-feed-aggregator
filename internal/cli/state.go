package cli

import (
	"github.com/G0SU19O2/rss-feed-aggregator/internal/config"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
)

type State struct {
	Cfg *config.Config
	Db  *database.Queries
}
