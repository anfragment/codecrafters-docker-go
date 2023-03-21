package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"syscall"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	command := os.Args[3]
	args := os.Args[4:len(os.Args)]

	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal("tmpDir: ", err)
	}
	cmdPath := path.Join(tmpDir, command)

	if err := exec.Command("mkdir", "-p", path.Dir(cmdPath)).Run(); err != nil {
		log.Fatal("mkdir: ", err)
	}
	if err := exec.Command("cp", "-f", command, cmdPath).Run(); err != nil {
		log.Fatal("cp: ", err)
	}

	syscall.Chroot(tmpDir)
	os.Chdir("/")
	cmd := exec.Command(command, args...)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		} else {
			log.Fatal(err)
		}
	}
}
