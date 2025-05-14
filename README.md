# Gator RSS Feed Aggregator

Gator is a command-line RSS feed aggregator written in Go. It allows you to register, follow feeds, aggregate posts, and browse them from your terminal.

## Prerequisites

- **MySQL**: You need a running MySQL server for Gator to store users, feeds, and posts.
- **Go**: Version 1.18 or newer is recommended. [Download Go here](https://golang.org/dl/).

## Installation

Install the Gator CLI using `go install`:

```sh
go install github.com/G0SU19O2/rss-feed-aggregator@latest
```

This will install the `rss-feed-aggregator` (or `gator`) binary to your `$GOPATH/bin` or `$HOME/go/bin`.

## Configuration

Before running the program, you need to set up a config file in your home directory. The config file is named `.gatorconfig.json` and should look like this:

```json
{
  "db_url": "user:password@tcp(127.0.0.1:3306)/gator_db?parseTime=true",
  "current_user_name": ""
}
```

- Replace `user`, `password`, and `gator_db` with your MySQL credentials and database name.

## Database Migration

Run your SQL migrations to set up the database schema. You can use [goose](https://github.com/pressly/goose) or your preferred migration tool:

```sh
goose -dir sql/schema mysql "user:password@/gator_db?parseTime=true" up
```

## Usage

Run the CLI from your terminal:

```sh
rss-feed-aggregator <command> [args...]
```

### Common Commands

- **register**  
  Register a new user.
  ```sh
  rss-feed-aggregator register <username>
  ```

- **login**  
  Log in as an existing user.
  ```sh
  rss-feed-aggregator login <username>
  ```

- **addfeed**  
  Add a new RSS feed to follow.
  ```sh
  rss-feed-aggregator addfeed "<Feed Name>" "<Feed URL>"
  ```

- **agg**  
  Aggregate (scrape) feeds at a given interval (e.g., every 10 minutes).
  ```sh
  rss-feed-aggregator agg 10m
  ```

- **browse**  
  Browse recent posts from feeds you follow. Optionally specify a limit (default is 2).
  ```sh
  rss-feed-aggregator browse 5
  ```

- **follow**  
  Follow a feed by its ID.
  ```sh
  rss-feed-aggregator follow <feed_id>
  ```

- **unfollow**  
  Unfollow a feed by its ID.
  ```sh
  rss-feed-aggregator unfollow <feed_id>
  ```

## Notes

- Make sure your MySQL server is running and accessible.
- The config file must be present in your home directory for the CLI to work.
- For development, you can run the CLI directly with `go run main.go <command> [args...]`.

---

Enjoy aggregating your favorite RSS feeds with Gator!