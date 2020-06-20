package main

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func runContainer(config Config) {
	args1 := []string{
		"docker",
		"run",
		"-it",
		"--rm",
	}
	opts := strings.Split(config.RunOption, " ")
	args2 := make([]string, 0, len(config.Dirs)*2+len(config.Envs)*2)
	for _, dir := range config.Dirs {
		args2 = append(args2, "--mount")
		args2 = append(args2, "source="+dir.Volume+",target=/work/"+dir.Name)
	}
	for _, e := range config.Envs {
		args2 = append(args2, "--env")
		args2 = append(args2, e)
	}
	args := append(args1, opts...)
	args = append(args, args2...)
	args = append(args, config.Img)

	bin, err := exec.LookPath("docker")
	if err != nil {
		panic(err)
	}
	if err := syscall.Exec(bin, args, os.Environ()); err != nil {
		panic(err)
	}
}
