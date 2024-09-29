
# Gator CLI

Gator CLI is a command-line tool designed to manage RSS feeds. This documentation will guide you through the setup, configuration, and usage of Gator CLI.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- **PostgreSQL**: Gator uses PostgreSQL as its database.
- **Go**: You need Go installed to compile the application.

You will also need Docker for setting up PostgreSQL using Docker Compose.

## Installation

To install the `gator` CLI tool, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/github-username/repo-name.git
cd repo-name
```

2. Set up PostgreSQL using Docker Compose:

First, ensure you have Docker installed. Then, run the following command to start PostgreSQL:
```bash
docker-compose up -d
```

3. Set up the configuration file:

You need to create a configuration file in your home directory named `.gatorconfig.json`. The content of the file should be as follows: 
```json
{
  "db_url": "postgres://myuser:postgres@localhost:5432/gator?sslmode=disable"
}
```

If you modify the Docker Compose settings (like the port, user, or password), make sure to update the db_url accordingly.


4. Run the migrations using Goose:

Once PostgreSQL is running, you need to apply the database migrations:

```bash
cd sql/schema
goose postgres "postgres://myuser:postgres@localhost:5432/gator?sslmode=disable" up
```

5. Install the Gator CLI tool:

After the setup, you can install the gator binary:

```bash
go install
```

This will compile and install the binary globally. You can now use the gator command from anywhere on your system.


### Usage

Gator CLI comes with several subcommands to manage RSS feeds and users.

Here are some example commands:

Register a new user:

```bash
gator register username
```

Login as a user:
```bash
gator login username
```

Add an RSS feed (Requires user to be logged in):
```bash
gator addfeed "https://example.com/rss"
```

Follow a feed (Requires user to be logged in):
```bash
gator follow feedID
```

Browse RSS posts (Requires user to be logged in):
```bash
gator browse
```


By default, it will display the latest 2 posts. You can customize the limit:
```bash
gator browse --limit=5
```

### Commands
Here’s a full list of subcommands available in Gator:

- `login`: Logs in a user.
- `register`: Registers a new user.
- `reset`: Resets a user's password.
- `users`: Lists all users.
- `agg`: Aggregates RSS feeds.
- `addfeed`: Adds a new feed to follow.
- `feeds`: Lists all available feeds.
- `follow`: Follows a feed (requires login).
- `following`: Shows feeds the user is following (requires login).
- `unfollow`: Unfollows a feed (requires login).
- `browse`: Browses RSS posts (requires login).

### Development Mode

For development purposes, you can run the program directly using:

```bash
go run .
```
Example:

```bash
go run . register username
go run . login username
```

### Production Mode
Once installed with `go install`, you can run the binary directly:

```bash
gator register username
gator login username
```


### Stopping and Restarting PostgreSQL
To stop the PostgreSQL instance:

```bash
docker-compose down
```

To restart it:

```bash
docker-compose up -d
```

