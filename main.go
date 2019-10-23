//go:generate go run github.com/UnnoTed/fileb0x b0x.yaml

package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
)

func main() {
	fmt.Println()
	defer fmt.Println()

	var app = kingpin.New("AriaNg-go", "AriaNg in go http server")
	var address = app.Flag("address", "Bind address.").Short('a').Default(":18080").String()
	var serverCmd = app.Command("server", "Start http server for AriaNg").Alias("s")

	app.Version(buildVersion())
	app.VersionFlag.Short('v')
	app.HelpFlag.Short('h')

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch cmd {
	case serverCmd.FullCommand():
		startServer(*address)
	}
}
