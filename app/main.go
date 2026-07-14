package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)
	builtins := []string{"echo", "exit", "type"}

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		tokens := strings.Fields(input)
		cmd, args := tokens[0], tokens[1:]

		if cmd == "echo" {
			fmt.Println(strings.Join(args, " "))
			continue
		}

		if cmd == "exit" {
			break
		}

		if cmd == "type" {
			if slices.Contains(builtins, tokens[1]) {
				fmt.Println(tokens[1] + " is a shell builtin")
			} else if path, err := exec.LookPath(args[0]); err == nil {
				fmt.Println(args[0] + " is " + path)
			} else if tokens[0] == "type" {
				fmt.Println(tokens[1] + " not found")
			}
			continue
		}

		fmt.Println(cmd + ": command not found")
	}
}
