package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func fzf(elements []string) string {
	result := withFilter("fzf -m", func(in io.WriteCloser) {
		for _, element := range elements {
			fmt.Fprintf(in, "%s\n", element)
		}
	})

	return result[0]
}

func withFilter(command string, input func(in io.WriteCloser)) []string {
	shell := os.Getenv("SHELL")
	if len(shell) == 0 {
		shell = "sh"
	}
	cmd := exec.Command(shell, "-c", command)
	cmd.Stderr = os.Stderr
	in, _ := cmd.StdinPipe()
	go func() {
		input(in)
		in.Close()
	}()
	result, _ := cmd.Output()
	return strings.Split(string(result), "\n")
}
