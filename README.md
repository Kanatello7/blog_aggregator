# Blog_aggregator - A Simple RSS Aggregator CLI

`blog_aggregator` is a command-line RSS feed aggregator written in Go as part of the Boot.dev "Learn Go" course. It allows you to manage users, add RSS feeds, scrape them periodically, and browse posts — all stored in a local PostgreSQL database.

## Features

- User registration and authentication
- Add, list, and delete RSS feeds
- Automatically scrape and store posts from feeds
- Browse posts with pagination
- Configurable scrape interval

## Prerequisites

You need the following installed on your machine:

- **Go** (version 1.22 or higher recommended): https://go.dev/dl/
- **PostgreSQL** (version 13 or higher): https://www.postgresql.org/download/

Make sure PostgreSQL is running and you have a database ready.

## Installation

Install the `blog_aggregator` CLI globally using `go install`:

```bash
go install github.com/Kanatello7/blog_aggregator@latest
```

This will place the `blog_aggregator` binary in your `$GOPATH/bin` (usually `~/go/bin`).  
Make sure `$GOPATH/bin` (or `$HOME/go/bin`) is in your `$PATH`.

Verify installation:

```bash
blog_aggregator
```

## Configuration

Create a config file named `gatorconfig.json` in your home directory (or any location you prefer):

```json
{
  "db_url": "postgres://yourusername:yourpassword@localhost:5432/gator",
  "current_user_name": ""
}
```

Replace the connection string with your actual PostgreSQL credentials and database name (`gator` is recommended).

You can place the file at:
- Default location: `~/.gatorconfig.json` (recommended)

The database will be automatically initialized the first time you run a command that needs it.

## Usage

### Common Commands

```bash
# Register a new user
blog_aggregator register username

# Log in as a user (sets current_user_name in config)
blog_aggregator login username

# Create the database schema (only needed once)
blog_aggregator reset   # caution: deletes everything
# or run any command — it will auto-migrate on first use

# Add an RSS feed
blog_aggregator addfeed "My Blog" "https://example.com/feed"

# List all feeds
blog_aggregator feeds

# Start scraping feeds in the background (default every 60 seconds)
blog_aggregator scrape

# Browse posts (default limit 2 posts, 1 per feed)
blog_aggregator browse
blog_aggregator browse 10   # show up to 10 posts per feed
```

### Example Session

```bash
blog_aggregator register alice
blog_aggregator login alice
blog_aggregator addfeed "Boot.dev Blog" "https://boot.dev/feed.xml"
blog_aggregator addfeed "RSS Specs" "https://www.rssboard.org/files/sample-rss-2.xml"
blog_aggregator feeds
blog_aggregator scrape   # press Ctrl+C to stop
blog_aggregator browse
```

## Available Commands
- `register` – create a new user
- `login`    – set the current user
- `users`    – list all users
- `addfeed`  – add a feed (requires a name and URL)
- `feeds`    – list all feeds
- `follow`   – follow a feed (as current user)
- `following`– list feeds you're following
- `unfollow` – stop following a feed
- `browse`   – read posts (with optional --limit)
- `scrape`   – continuously fetch new posts
- `reset`    – delete all data (useful for testing)

## Contributing

Feel free to open issues or PRs! This project was built for learning purposes.

Made with ❤️ while learning Go on [Boot.dev](https://boot.dev)
```

### 3. & 4. Installation & Setup Instructions (covered in README above)

All required explanations are clearly documented in the README:

- Need Go + PostgreSQL
- How to `go install`
- How to set up `gatorconfig.json`
- How to run common commands
