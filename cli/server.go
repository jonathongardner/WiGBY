package cli

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"net/http"

	"github.com/jonathongardner/wegyb/server"
	"github.com/jonathongardner/wegyb/camera"
	"github.com/urfave/cli/v2"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
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
		&cli.StringFlag{
			Name:    "vision-xml",
			Value:   "",
			Usage:   "Cascade Classifier car xml file",
			EnvVars: []string{"WEGYB_VISION_XML"},
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
		visionConfig := c.String("vision-xml")
		deviceId := c.Int("device")

		log.Info("Starting...")

		ctx, cancel := context.WithCancel(context.Background())

		// listen for ctrl + c and gracefully shutdown
		go func() {
			c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)

			<-c
			log.Info("Gracefully Shuting down...")
			cancel()
		}()
		g, gCtx := errgroup.WithContext(ctx)

		// setup camera and start capturing frames in another routine
		ch := camera.NewHub()
		vision := camera.NewVision(visionConfig)

		g.Go(func() error {
			return ch.Run(deviceId, vision, gCtx)
		})

		// create server, and start in another thread with another thread lstening for closing
		httpServer := server.NewServer(hostPort, ch, ui)
		g.Go(func() error {
			err := httpServer.ListenAndServe()
			if err != http.ErrServerClosed {
				return err
			}
			return nil
		})
		g.Go(func() error {
			<- gCtx.Done()
			return httpServer.Shutdown(context.Background())
		})

		// now wait to see if any errors are raised, if one is raised than it will call cancel and end
		if err := g.Wait(); err != nil {
			log.Errorf("exit reason: %s \n", err)
		}

		log.Info("...Closing")

		return nil
	},
}
