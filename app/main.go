package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	command := os.Args[3]
	args := os.Args[4:len(os.Args)]

	cmd := exec.Command(command, args...)

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stderrPipe.Close()
	go func() {
		io.Copy(os.Stderr, stderrPipe)
	}()

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(os.Stdout, string(output))
}
