package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/razorpay/MachineRound/internal/config"
	"github.com/razorpay/MachineRound/internal/constants"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/pressly/goose"

	// Init DB drivers.

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flags.String(constants.MigrationDir, constants.DefaultMigrationDir, "directory with migration files")
	env   = flags.String(constants.Env, constants.Development, "Application env : prod/dev")
)

const (
	POSTGRES   = "postgres"
	CREATE     = "create"
	HELP       = "--help"
	SHORT_HELP = "-h"
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])

	args := flags.Args()

	if len(args) > 1 && args[0] == CREATE {
		if err := goose.Run(CREATE, nil, *dir, args[1:]...); err != nil {
			logrus.Fatalf("goose run: %v", err)
		}

		return
	}

	if len(args) < 1 {
		flags.Usage()

		return
	}

	if args[0] == HELP || args[0] == SHORT_HELP {
		flags.Usage()

		return
	}

	config.LoadConfig(constants.DefaultBasePath, *env)
	command := args[0]
	driver := POSTGRES

	var databaseConfig config.Database

	databaseConfig = config.GetConfig().Database

	dbstring := databaseConfig.URL()
	if err := goose.SetDialect(driver); err != nil {
		logrus.Fatalf("goose run: %v", err)
	}

	db, err := sql.Open(driver, dbstring)

	if err != nil {
		logrus.Fatalf("-dbstring=%q: %v\n", dbstring, err)
	}

	arguments := []string{}

	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}
	fmt.Println(command, db, arguments)
	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		logrus.Fatalf("goose run: %v", err)
	}
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}

var (
	usagePrefix = `Usage: goose [OPTIONS] DRIVER DBSTRING COMMAND
Drivers:
    postgres
    mysql
    sqlite3
    redshift
Examples:
    goose sqlite3 ./foo.db status
    goose sqlite3 ./foo.db create init sql
    goose sqlite3 ./foo.db create add_some_column sql
    goose sqlite3 ./foo.db create fetch_user_data go
    goose sqlite3 ./foo.db up
    goose postgres "user=postgres dbname=postgres sslmode=disable" status
    goose mysql "user:password@/dbname?parseTime=true" status
    goose redshift "postgres://user:password@qwerty.us-east-1.redshift.amazonaws.com:5439/db" status
Options:
`

	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version
`
)
