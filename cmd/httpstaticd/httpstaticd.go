package main

import (
	"os"

	"github.com/flaccid/httpstaticd"
	"github.com/urfave/cli"

	log "github.com/sirupsen/logrus"
)

var (
	VERSION = "v0.0.0-dev"
)

func beforeApp(c *cli.Context) error {
	if c.GlobalBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "httpstaticd"
	app.Version = VERSION
	app.Usage = "basic http server for static files"
	app.Action = start
	app.Before = beforeApp
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "directory,d",
			Usage: "directory of static files to host",
			Value: "./",
		},
		cli.BoolFlag{
			Name:  "listings,l",
			Usage: "enable directory listings",
		},
		cli.IntFlag{
			Name:  "port,p",
			Usage: "port the listen on",
			Value: 8080,
		},
		// todo: always on, need to implement
		cli.BoolFlag{
			Name:  "access-log,a",
			Usage: "enable access logging of requests",
		},
		cli.BoolFlag{
			Name:  "debug,D",
			Usage: "run in debug mode",
		},
	}
	app.Run(os.Args)
}

func start(c *cli.Context) error {
	httpstaticd.Serve(c.String("directory"), c.Int("port"), c.Bool("listings"), c.Bool("access-log"))

	return nil
}
