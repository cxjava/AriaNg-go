//go:generate packr2

package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func init() {
	cli.VersionPrinter = func(c *cli.Context) {
		showVersion()
	}
}

var address string

func main() {
	app := cli.NewApp()
	app.Name = "AriaNg-go"
	app.Usage = "AriaNg in go http server"
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "address, a",
			Value:       ":18080",
			Usage:       "bind address",
			Destination: &address,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Start server for AriaNg",
			// Flags: []cli.Flag{
			// 	cli.StringFlag{
			// 		Name:  "path, p",
			// 		Usage: "template path",
			// 	},
			// },
			Action: func(c *cli.Context) error {
				startServer(address)
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Failed to run app with %s: %v", os.Args, err)
	}
}
