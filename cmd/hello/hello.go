package hello

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

var Command = cli.Command{
	Name:      "hello",
	Aliases:   []string{"h"},
	ArgsUsage: "<user-name>",
	Usage:     "print hello user",
	Flags: []cli.Flag{
		flagQuoteName,
	},
	Action: action,
}

var flagQuoteName = cli.BoolFlag{
	Name:  "quote-name",
	Usage: "quote the name in the output",
}

func action(c *cli.Context) error {
	if c.NArg() == 0 {
		_ = cli.ShowCommandHelp(c, c.Command.Name)
		return errors.New("missing required argument")
	}
	logger := log.With().Str("name", c.Command.Name).Logger()
	if c.Bool(flagQuoteName.Name) {
		logger.Debug().Msg("printing quoted")
		fmt.Fprintf(c.App.Writer, "Hello %q.\n", c.Args().First())
	} else {
		logger.Debug().Msg("printing un-quoted")
		fmt.Fprintf(c.App.Writer, "Hello %s.\n", c.Args().First())
	}
	return nil
}
