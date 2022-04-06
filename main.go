package main

import (
	"fmt"
	"log"
	"os"
	"embed"
	"io/fs"

	"github.com/jonathongardner/wegyb/wegyb"
	"github.com/urfave/cli/v2"
)

//go:embed public/*
var uiFolder embed.FS

// type enumValue struct {
// 	Enum     []string
// 	Default  string
// 	selected string
// }
//
// func (e *enumValue) Set(value string) error {
// 	for _, enum := range e.Enum {
// 		if enum == value {
// 			e.selected = value
// 			return nil
// 		}
// 	}
//
// 	return fmt.Errorf("allowed values are %s", strings.Join(e.Enum, ", "))
// }
//
// func (e enumValue) String() string {
// 	if e.selected == "" {
// 		return e.Default
// 	}
// 	return e.selected
// }
// &cli.StringFlag{
// 	Name:    "codec",
// 	Value:   &enumValue{
// 		Enum: []string{"mmal", "x264", "openh264", "vpx"},
// 		Default: "mmal"
// 	},
// 	Usage:   "Camera codec to use",
// 	EnvVars: []string{"WEGYB_CODEC"},
// },

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Version)
	}
	app := &cli.App{
		Name: "wegyb",
		Version: "0.0.1",
		Usage: "We got your back!",
		Commands: []*cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "start server",
			Flags: []cli.Flag {
				&cli.StringFlag{
					Name:    "host",
					Value:   "",
					Usage:   "Host to serve on (example: 0.0.0.0, localhost, etc)",
					EnvVars: []string{"WEGYB_HOST"},
				},
				&cli.StringFlag{
					Name:    "port",
					Value:   "3000",
					Usage:   "Port to publich to",
					EnvVars: []string{"WEGYB_PORT"},
				},
			},
			Action:  func(c *cli.Context) error {
				hostPort := c.String("host") + ":" + c.String("port")
				log.Println("Starting...")
				ui, _ := fs.Sub(uiFolder, "public")
				// Block forever
				wegyb.ListenAndServe(hostPort, ui)
				log.Println("...Closing")

				return nil
			},
		},
	},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
