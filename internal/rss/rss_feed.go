package rss

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	request, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "gator")
	res, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var feed RSSFeed
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feed.Channel.Item[i] = item
	}

	return &feed, nil
}

func ScrapeFeeds(ctx context.Context, db *database.Queries) error {
	dbFeed, err := db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("can't get any feed")
	}
	time := time.Now()
	err = db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{LastFetchedAt: sql.NullTime{Time: time, Valid: true}, UpdatedAt: time, ID: dbFeed.ID})
	if err != nil {
		return err
	}
	feed, err := FetchFeed(ctx, dbFeed.Url)
	if err != nil {
		return err
	}
	println("=========================")
	for _, item := range feed.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}
	return nil
}
