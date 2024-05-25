// This is custom goose binary with mysql support only.

package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/pressly/goose/v3"

	_ "github.com/goose-go-migrate-example/infrastructure/db/migrations"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", ".", "directory with migration files")
)

func main() {
	flags.Parse(os.Args[1:]) // Parse first to handle the flags.
	args := flags.Args()     // Then retrieve the remaining arguments.

	if len(args) < 2 { // Adjusting for the actual minimum arguments required.
		flags.Usage()
		return
	}

	dbstring, command := args[0], args[1] // Adjusted indices according to the new understanding.

	db, err := goose.OpenDBWithDriver("mysql", dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 2 {
		arguments = append(arguments, args[2:]...)
	}

	if err := goose.RunContext(context.Background(), command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
