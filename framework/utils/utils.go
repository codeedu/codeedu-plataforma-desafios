package utils

import (
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"os"
	"log"
	"strconv"
	"errors"
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

func DebugMode() bool {
	debug, err := Getenv("debug");

	if err != nil {
		return false
	}

	v, _:= strconv.ParseBool(debug)

	return v
}

func Printf(format string, v ...interface{}) {
	if DebugMode() {
		log.Printf(format, v...)
	}
}

func Getenv(v string) (string, error) {
	//Load environmenatal variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv(v)	

	if key == "" {
		return key, errors.New("Error loading variable " + v)
	}

	return key, nil
}
