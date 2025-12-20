package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/CodeZeroSugar/gator/internal/config"
	"github.com/CodeZeroSugar/gator/internal/database"
	_ "github.com/lib/pq"
)

const dbURL = "postgres://postgres:password@localhost:5432/gator?sslmode=disable"

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	cfgState := &state{
		configPtr: &cfg,
		db:        dbQueries,
		dbPool:    db,
	}
	cfgState.configPtr.DBURL = dbURL

	commands := commands{commandMap: make(map[string]func(*state, command) error)}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", reset)
	commands.register("users", users)
	commands.register("agg", agg)
	commands.register("addfeed", addFeed)
	commands.register("feeds", feedsHandler)
	commands.register("follow", follow)
	commands.register("following", following)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("error: need at least two args")
	}

	currentCommand := command{
		name: args[1],
		args: args[2:],
	}

	if err := commands.run(cfgState, currentCommand); err != nil {
		log.Fatal(err)
	}

	/*contents, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", contents)
	*/
}
