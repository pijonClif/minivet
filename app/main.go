package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)
	builtins := []string{"echo", "exit", "type"}

	for {
		fmt.Print("$ ")

		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		cmd = strings.TrimSpace(cmd)
		tokens := strings.Split(cmd, " ")

		if tokens[0] == "type" && slices.Contains(builtins, tokens[1]) {
			fmt.Println(tokens[1] + " is a shell builtin")
		} else if tokens[0] == "type" {
			fmt.Println(tokens[1] + " not found")
		} else if cmd == "exit" {
			break
		} else if strings.HasPrefix(cmd, "echo ") {
			fmt.Println(cmd[5:])
		} else {
			fmt.Println(cmd + ": command not found")
		}

	}
}
