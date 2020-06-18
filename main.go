package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tedicon [flag] # run vim container refer to ~/.edicon.json\n")
	flag.PrintDefaults()
}

func main() {
  home, err := homedir.Dir()
  if err != nil {
    panic(err)
  }
  configFilePath := flag.String("f", home + "/.edicon.json", "config file path for run container")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 0 {
		usage()
		os.Exit(2)
	}

	// parse json
  c, err := parseConfigFile(*configFilePath)
  if err != nil {
    panic(err)
  }

  runContainer(c)
}
