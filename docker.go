package main

import (
	"os"
	"os/exec"
	"syscall"
)

func runContainer(img string, dirs []Dir){
	args1 := []string{
    "docker",
		"run",
		"-it",
    "--rm",
  }
  args3 := []string{
     "--mount",
		"source=vim,target=/root",
		"--name",
		"vim",
		img,
	}
  args2 := make([]string, 0, len(dirs)*2)
  for _, dir := range dirs {
    args2 = append(args2, "--mount")
    args2 = append(args2, "source=" + dir.Volume + ",target=/work" )
  }
  args := append(args1, args2...)
  args = append(args, args3...)

  bin, err := exec.LookPath("docker")
  if err != nil {
    panic(err)
  }
  syscall.Exec(bin, args, os.Environ())
}
