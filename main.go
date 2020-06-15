package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	configFilePath = flag.String("f", "~/.convim.json", "config file path for run container")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tconvim [flag] # run vim container refer to ~/.convim.json\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 0 {
		usage()
		os.Exit(2)
	}

	// parse json

	// ran docker container
}
