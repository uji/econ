package main

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func runContainer(config Config) {
	cmd := []string{
		"docker",
		"run",
		"-it",
		"--rm",
	}
	runOpts := strings.Split(config.RunOption, " ")
	mountOpts := make([]string, 0, len(config.Dirs)*2+len(config.Envs)*2)
	for _, dir := range config.Dirs {
		mountOpts = append(mountOpts, "--mount")
		mountOpts = append(mountOpts, "source="+dir.Volume+",target=/work/"+dir.Name)
	}
	for _, e := range config.Envs {
		mountOpts = append(mountOpts, "--env")
		mountOpts = append(mountOpts, e)
	}

	args := make([]string, 0, 5+len(runOpts)+len(config.Dirs)*2+len(config.Envs)*2)
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
