package main

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli"
	"log"
	"opentodo/db"
	"opentodo/utils"
	"os"
	"time"
)

func main() {
	const MigrationPath = "file://db/migrations"

	app := cli.NewApp()
	dbConfig := db.GetDBConfig()
	connection, err := db.Connect(dbConfig)
	driver, err := postgres.WithInstance(connection.DB(), &postgres.Config{})

	_, err = migrate.NewWithDatabaseInstance(MigrationPath, dbConfig.DBName, driver)

	utils.Must(err)

	app.Commands = []cli.Command{
		{
			Name:        "create",
			Description: "Create a new migration",
			Action: func(context *cli.Context) {
				migrationName := context.Args().Get(0)
				_, err := migrate.NewMigration(nil, migrationName, uint(time.Now().Second()), time.Now().Second()+1)

				utils.Must(err)
			},
		},
	}

	err = app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
