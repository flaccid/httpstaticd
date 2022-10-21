package main

import (
	"os"

	"github.com/flaccid/httpstaticd"
	"github.com/urfave/cli"

	log "github.com/sirupsen/logrus"
)

func beforeApp(c *cli.Context) error {
	switch c.GlobalString("log-level") {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "":
		log.SetLevel(log.InfoLevel)
	default:
		log.Fatalf("%s is an invalid log level", c.GlobalString("log-level"))
	}

	log.Info("using log level " + log.GetLevel().String())

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "httpstaticd"
	app.Version = httpstaticd.VERSION
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
			Name:  "listings,L",
			Usage: "enable directory listings",
		},
		cli.IntFlag{
			Name:  "port,p",
			Usage: "port the listen on",
			Value: 8080,
		},
		cli.BoolFlag{
			Name:  "cors,c",
			Usage: "enable cors (not supported with directory listings)",
		},
		cli.BoolFlag{
			Name:  "access-log,a",
			Usage: "enable access logging of requests",
		},
		cli.BoolFlag{
			Name: "no-log-healthz",
			Usage: "do not log health check requests",
		},
		cli.StringFlag{
			Name:  "log-level,l",
			Usage: "log level to use (debug,warn,error,info); default: info",
		},
	}
	app.Run(os.Args)
}

func start(c *cli.Context) error {
	httpstaticd.Serve(c.String("directory"),
		c.Bool("listings"),
		c.Int("port"),
		c.Bool("cors"),
		c.Bool("access-log"))

	return nil
}
