package cli

import (
	"fmt"
	"os"
	"io/fs"

	"github.com/urfave/cli/v2"
)

var ui fs.FS

func Run(uiPassed fs.FS) (error) {
	ui = uiPassed

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Version)
	}

	app := &cli.App{
		Name: "wegyb",
		Version: "0.0.1",
		Usage: "We got your back!",
		Commands: []*cli.Command{
  		serverCommand,
  	},
	}
	return app.Run(os.Args)
}
