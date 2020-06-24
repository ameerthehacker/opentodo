package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:        "create",
			Description: "Create a new migration",
			Action: func(context *cli.Context) {
				migrationName := context.Args().Get(0)

				fmt.Println(migrationName)
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
