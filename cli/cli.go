package cli

import (
	"fmt"
	"os"
	"io/fs"
	"github.com/jonathongardner/wegyb/app"

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
		Version: app.Version,
		Usage: "We got your back!",
		Commands: []*cli.Command{
  		serverCommand,
  	},
	}
	return app.Run(os.Args)
}
