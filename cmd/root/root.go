package root

import (
	"os"

	"github.com/Eun/go-bin-template/cmd/hello"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

var (
	name        string
	version     string
	commit      string
	date        string
	description string
	author      string
)

var cliFlagLogFile = cli.StringFlag{
	Name:   "logfile",
	Usage:  "Logfile to write to",
	EnvVar: "LOGFILE",
	Value:  "",
}
var cliFlagDebug = cli.BoolFlag{
	Name:   "debug",
	Usage:  "Enable debug log",
	EnvVar: "DEBUG",
}

func Run() int {
	var logfile *os.File
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	app := cli.NewApp()
	app.Name = name
	app.Version = version + " " + commit + " " + date
	app.Description = description
	app.Author = author
	app.Flags = []cli.Flag{
		cliFlagLogFile,
		cliFlagDebug,
	}
	app.Commands = []cli.Command{
		hello.Command,
	}
	app.Before = func(context *cli.Context) error {
		loglevel := zerolog.InfoLevel
		if context.Bool(cliFlagDebug.Name) {
			loglevel = zerolog.DebugLevel
		}
		log.Logger = log.Level(loglevel)

		logfilePath := context.String(cliFlagLogFile.Name)
		if logfilePath != "" {
			var err error
			//nolint:gomnd // allow magic number 0666
			logfile, err = os.OpenFile(logfilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0666)
			if err != nil {
				return errors.Wrap(err, "unable to create logfile")
			}
			log.Logger = log.Output(logfile)
		}
		log.Debug().Str("logfile", logfilePath).Msg("debug log enabled")
		return nil
	}

	defer func() {
		if logfile == nil {
			return
		}
		log.Logger = log.Output(os.Stderr)
		if err := logfile.Close(); err != nil {
			log.Fatal().Err(err).Msg("unable to close logfile")
		}
	}()

	if err := app.Run(os.Args); err != nil {
		log.Error().Err(err).Msg("error during execution")
		return 1
	}
	return 0
}
