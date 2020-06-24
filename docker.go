package main

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

const (
	workDir = "/econ"
)

func runContainer(config Config, volume string) {
	cmd := []string{
		"docker",
		"run",
		"-it",
		"--rm",
		"--workdir=" + workDir,
		"--mount",
		"source=" + volume + ",target=" + workDir,
	}
	runOpts := strings.Split(config.RunOption, " ")
	mountOpts := make([]string, 0, 2+len(config.Envs)*2)
	for _, e := range config.Envs {
		mountOpts = append(mountOpts, "--env")
		mountOpts = append(mountOpts, e)
	}

	args := make([]string, 0, len(cmd)+len(runOpts)+len(config.Envs)*2+1)
	args = append(args, cmd...)
	if config.RunOption != "" {
		args = append(args, runOpts...)
	}
	args = append(args, mountOpts...)
	args = append(args, config.Img)

	bin, err := exec.LookPath("docker")
	if err != nil {
		panic(err)
	}
	if err := syscall.Exec(bin, args, os.Environ()); err != nil {
		panic(err)
	}
}

func isVolume(name string) bool {
	// TODO: implement checking volume existence
	return true
}
