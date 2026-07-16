package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var builtins = []string{"echo", "exit", "type", "pwd", "cd"}

var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		tokens := strings.Fields(input)
		if len(tokens) == 0 {
			continue
		}
		cmd, args := tokens[0], tokens[1:] //user_command, user_command_arguments

		switch cmd {
		case "echo":
			fmt.Println(strings.Join(args, " "))
			continue

		case "exit":
			os.Exit(0)

		case "pwd":
			handlePwd()
			continue

		case "cd":
			handleCd(args[0])
			continue

		case "type":
			handleType(args[0])
			continue
		}

		if _, err := exec.LookPath(cmd); err == nil {
			prog := exec.Command(cmd, args...)
			prog.Stdout = os.Stdout
			prog.Stderr = os.Stderr
			prog.Run()
		} else {
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}

func handlePwd() {
	abs_dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(abs_dir)
}

func handleCd(arg string) {
	targetDir := arg
	if targetDir == "~" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("something evil has occurred [couldnt locate home path]")
		}
		targetDir = home
	}

	err := os.Chdir(targetDir)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", arg)
	}
}

func handleType(arg string) {
	if slices.Contains(builtins, arg) {
		fmt.Println(arg + " is a shell builtin")
	} else if path, err := exec.LookPath(arg); err == nil {
		fmt.Println(arg + " is " + path)
	} else {
		fmt.Println(arg + " not found")
	}
}
