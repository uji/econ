package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
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

  runContainer("vim", nil)
}

type mountDir struct {
  name string
  volume string
}

func runContainer(img string, dir []mountDir){
	args := []string{
    "docker",
		"run",
		"-it",
    "--rm",
		// "--mount",
		// "source=$(VOLUME),target=/work",
		"--mount",
		"source=vim,target=/root",
		"--name",
		"vim",
		img,
	}
  bin, err := exec.LookPath("docker")
  if err != nil {
    panic(err)
  }
  syscall.Exec(bin, args, os.Environ())
}
