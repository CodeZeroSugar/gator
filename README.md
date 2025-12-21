# Gator feed aggregator

Welcome to Gator, a tool that collects feeds and stores them in a database as posts for easy retrieval.

## Installation
Prerequisites:
- go 1.25.5
- psql (PostgreSQL) 18.1 

Install from the root directory of the program:
```bash 
go install
``` 
### Setup
1. Create the config file in the home directory ~/.gatorconfig.json: 
  - Manually input your connection string for the 'db_url' value, 'current_user_name' will be populated when the first user is created.
```json 
{
  "db_url":"postgres://postgres:<yourpassword>@localhost:5432/gator?sslmode=disable",
  "current_user_name":"<provided by program>"
}
```
2. run the program from the root of the directory with `./gator` or add it to your PATH to run as just `gator`.
```bash
export PATH=$PATH:/path/to/directory
```
#### Quick Start
1. Register a user
  - `gator register newUser`
2. Add a feed
  - `gator addfeed "Hacker News RSS" "https://hnrss.org/newest"`
3. Collect posts from the feed every 30 seconds
  - `gator agg 30s`
4. Browse the last 5 posts you collected
  - `gator browse 5`


##### Usage
1. `register` adds a new user to the database and sets them as the 'Current User'.
  - `gator register gatorUser`
2. `login` "logs into" an existing user and sets them as the 'Current User'.
  - `gator login gatorUser`
3. `addfeed` adds a new feed to the database. It requires the name of the feed and the URL as arguments.
  - `gator addfeed "Hacker News RSS" "https://hnrss.org/newest"`
4. `feeds` displays all added feeds and what user created the feed.
  - `gator feeds`
5. `follow` adds a feed to the current user's feed follow record. `unfollow` will remove it from their record.
  - `gator follow "https://hnrss.org/newest"`
  - `gator unfollow "https://hnrss.org/newest"`
6. `following` displays the followed feeds for the current user.
  - `gator following`
7. `agg` collects posts from the current user's followed feeds. It requires the amount of time between requests as an argument.
  - `gator agg 30s`
  - This command is intended to be run in the background on a separate terminal session.
8. `browse` lets the current user view all the posts from their followed feeds. This command takes an optional 'limit' argument to change the number of posts displayed. The default is 2.
  - `gator browse`
  - `gator browse 5`

Thank you for checking out this project!

