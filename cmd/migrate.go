package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"opentodo/db"
	"opentodo/utils"
	"os"
	"strconv"
	"time"
)

func main() {
	const MigrationPath = "db/migrations"

	err := godotenv.Load()
	utils.Warn(err, "No .env file was found in the root, defaulting to use environment variables")

	app := cli.NewApp()
	dbConfig := db.GetDBConfig()
	connection, err := db.Connect(dbConfig)
	defer func() {
		if err != nil {
			connection.Close()
		}
	}()

	utils.Must(err)

	driver, err := postgres.WithInstance(connection.DB(), &postgres.Config{})

	utils.Must(err)

	migration, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", MigrationPath), dbConfig.DBName, driver)

	utils.Must(err)

	app.Commands = []cli.Command{
		{
			Name:        "create",
			Description: "Create a new migration",
			Action: func(ctx *cli.Context) {
				migrationName := ctx.Args().Get(0)
				upMigrationFilename := fmt.Sprintf("%s/%d_%s.up.sql", MigrationPath, time.Now().Unix(), migrationName)
				downMigrationFilename := fmt.Sprintf("%s/%d_%s.down.sql", MigrationPath, time.Now().Unix(), migrationName)

				// don't overwrite existing files
				if _, err = os.Stat(upMigrationFilename); !os.IsNotExist(err) {
					log.Panicf("%s already exists", upMigrationFilename)
				}
				if _, err = os.Stat(downMigrationFilename); !os.IsNotExist(err) {
					log.Panicf("%s already exists", downMigrationFilename)
				}

				err = ioutil.WriteFile(upMigrationFilename, []byte{}, 0664)
				err = ioutil.WriteFile(downMigrationFilename, []byte{}, 0664)

				utils.Must(err)

				fmt.Printf("Created %s, %s", upMigrationFilename, downMigrationFilename)
			},
		}, {
			Name:        "up",
			Description: "Run pending migrations if any",
			Action: func(ctx *cli.Context) {
				err = migration.Up()

				utils.Must(err)

				fmt.Println("Migration done!")
			},
		},
		{
			Name:        "down",
			Description: "Reverse the last migration",
			Action: func(ctx *cli.Context) {
				err = migration.Down()

				utils.Must(err)

				fmt.Println("Migration done!")
			},
		},
		{
			Name:        "force",
			Description: "Force a particular migration version",
			Action: func(ctx *cli.Context) {
				if ctx.NArg() > 0 {
					version := ctx.Args().Get(0)
					versionNumber, err := strconv.ParseInt(version, 10, 64)

					if err != nil {
						log.Panicln("Invalid version number", version)
					}

					err = migration.Force(int(versionNumber))

					utils.Must(err)

					fmt.Println("Migration done!")
				} else {
					fmt.Println("No version number was provided")
				}
			},
		},
	}

	err = app.Run(os.Args)

	utils.Must(err)
}
