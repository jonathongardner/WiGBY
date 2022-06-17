package cli

import (
	"github.com/jonathongardner/wegyb/server"
	"github.com/urfave/cli/v2"
	log "github.com/sirupsen/logrus"
)

var serverCommand =  &cli.Command{
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
			Aliases: []string{"p"},
			Value:   "3000",
			Usage:   "Port to publich to",
			EnvVars: []string{"WEGYB_PORT"},
		},
		&cli.IntFlag{
			Name:    "device",
			Aliases: []string{"d"},
			Value:   0,
			Usage:   "Camera device to use",
			EnvVars: []string{"WEGYB_DEVICE"},
		},
	},
	Action:  func(c *cli.Context) error {
		hostPort := c.String("host") + ":" + c.String("port")
		log.Info("Starting...")
		// Block forever
		server.ListenAndServe(hostPort, c.Int("device"), ui)
		log.Info("...Closing")

		return nil
	},
}
