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
		if len(tokens)==0{
			continue
		}
		cmd, args := tokens[0], tokens[1:] //user_command, user_command_arguments

		if cmd == "echo" {
			fmt.Println(strings.Join(args, " "))
			continue
		}

		if cmd == "exit" {
			break
		}

		if cmd == "type" {
			if slices.Contains(builtins, tokens[1]) {
				fmt.Println(args[0] + " is a shell builtin")
			} else if path, err := exec.LookPath(args[0]); err == nil {
				fmt.Println(args[0] + " is " + path)
			} else if cmd == "type" {
				fmt.Println(args[0] + " not found")
			}
			continue
		} else if _, err := exec.LookPath(cmd); err == nil {
			prog := exec.Command(cmd, args...)
			prog.Stdout = os.Stdout
			prog.Stderr = os.Stderr
			prog.Run()
		} else {
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
