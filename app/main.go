package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		cmd = strings.TrimSpace(cmd)
		if cmd == "exit" {
			break
		} else if strings.HasPrefix(cmd, "echo ") {
			fmt.Println(cmd[5:])
		} else {
			fmt.Println(cmd + ": command not found")
		}

	}
}
