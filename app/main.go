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
	builtins := []string{"echo", "exit", "type", "pwd", "cd"}

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


		switch cmd{
		case "echo":
			fmt.Println(strings.Join(args, " "))
			continue

		case "exit":
			os.Exit(0)

		case "pwd":
			abs_dir, err:= os.Getwd()
			if err!=nil{
				fmt.Println(err)
			}
			fmt.Println(abs_dir)
			continue
		
		case "cd":
			targetDir:=args[0]

			if targetDir=="~"{
				home,err:=os.UserHomeDir()
				if err!=nil{
					fmt.Printf("something evil has occurred [couldnt locate home path]")
				}
				targetDir=home 
			}
			
			err:=os.Chdir(targetDir)

			if err!=nil{
				fmt.Printf("%s: %s: No such file or directory\n",cmd, args[0])
			}
			continue
		}

		if cmd == "type" {
			if slices.Contains(builtins, args[0]) {
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
