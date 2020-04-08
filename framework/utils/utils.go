package utils

import (
	"github.com/jessevdk/go-flags"
	"os"
)

type ConfigOptions struct {
	Token string `short:"t" long:"token" description:"Token" required:"true"`
}

func CheckConfigFlags(parser *flags.Parser) {
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
