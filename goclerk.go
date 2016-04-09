package main

import (
	"os"

	"github.com/jonaswouters/goclerk/cmd"

	"github.com/codegangsta/cli"
)

// AppVersion Version of the ap
const AppVersion = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "GoClerk"
	app.Usage = "Accounting & More"
	app.Version = AppVersion
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		cmd.Web,
		cmd.Setup,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
