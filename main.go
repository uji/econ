package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tecon [flag] [volume] # run vim container refer to ~/.econ.json\n")
	flag.PrintDefaults()
}

func main() {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	configFilePath := flag.String("f", home+"/.econ.json", "config file path for run container")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
		os.Exit(2)
	}

	// parse json
	c, err := parseConfigFile(*configFilePath)
	if err != nil {
		panic(err)
	}

	if !isVolume(flag.Arg(0)) {
		panic(errors.New("docker volume is not found"))
	}

	runContainer(c, flag.Arg(0))
}
