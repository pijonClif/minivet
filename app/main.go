package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

var builtins = []string{"echo", "exit", "type", "pwd", "cd", "history"}
var history []string

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

		tokens := tokenize(input)
		history = append(history, strings.TrimRight(input, "\n"))
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

		case "history":
			handleHistory(args)
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

func handleType(arg string) {
	if slices.Contains(builtins, arg) {
		fmt.Println(arg + " is a shell builtin")
	} else if path, err := exec.LookPath(arg); err == nil {
		fmt.Println(arg + " is " + path)
	} else {
		fmt.Println(arg + " not found")
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

func handleHistory(args []string) {
	start := 0
	if len(args) > 0 {
		if n, err := strconv.Atoi(args[0]); err == nil && n < len(history) {
			start = len(history) - n
		}
	}
	for i := start; i < len(history); i++ {
		fmt.Printf("%5d  %s\n", i+1, history[i])
	}
}

func tokenize(input string) []string {
	var tokens []string
	var curr strings.Builder
	started := false
	inSingle := false
	inDouble := false

	runes := []rune(input)
	n := len(runes)

	for i := 0; i < n; i++ {
		c := runes[i]

		switch {
		case inSingle:
			if c == '\'' {
				inSingle = false
			} else {
				curr.WriteRune(c)
			}
		case inDouble:
			if c == '"' {
				inDouble = false
			} else if c == '\\' && i+1 < n && strings.ContainsRune(`"\$`+"`", runes[i+1]) {
				i++
				curr.WriteRune(runes[i])
			} else {
				curr.WriteRune(c)
			}
		case c == '\\':
			started = true
			if i+1 < n {
				i++
				curr.WriteRune(runes[i])
			}
		case c == '\'':
			inSingle = true
			started = true
		case c == '"':
			inDouble = true
			started = true
		case c == ' ' || c == '\t' || c == '\n' || c == '\r':
			if started {
				tokens = append(tokens, curr.String())
				curr.Reset()
				started = false
			}
		default:
			started = true
			curr.WriteRune(c)
		}
	}

	if started {
		tokens = append(tokens, curr.String())
	}
	return tokens
}
