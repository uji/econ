package main

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func runContainer(config Config, volume string) {
	cmd := []string{
		"docker",
		"run",
		"-it",
		"--rm",
	}
	runOpts := strings.Split(config.RunOption, " ")
	mountOpts := make([]string, 0, 2+len(config.Envs)*2)
	mountOpts = append(mountOpts, "--mount")
	mountOpts = append(mountOpts, "source="+volume+",target=/work")
	for _, e := range config.Envs {
		mountOpts = append(mountOpts, "--env")
		mountOpts = append(mountOpts, e)
	}

	args := make([]string, 0, 5+len(runOpts)+2+len(config.Envs)*2)
	args = append(args, cmd...)
	args = append(args, runOpts...)
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
