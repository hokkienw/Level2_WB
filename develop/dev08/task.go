package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)

		handleCommand(input)

		if input == "exit" {
			break
		}
	}
}

func handleCommand(input string) {
	args := strings.Fields(input)
	if len(args) == 0 {
		return
	}

	command := args[0]
	commandArgs := args[1:]

	switch command {
	case "cd":
		if len(commandArgs) == 0 {
			fmt.Println("Usage: cd <directory>")
			return
		}
		err := os.Chdir(commandArgs[0])
		if err != nil {
			fmt.Println("Error changing directory:", err)
		}
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
		}
		fmt.Println(dir)
	case "echo":
		fmt.Println(strings.Join(commandArgs, " "))
	case "kill":
		if len(commandArgs) == 0 {
			fmt.Println("Usage: kill <process_id>")
			return
		}
		pid := commandArgs[0]
		err := exec.Command("kill", pid).Run()
		if err != nil {
			fmt.Println("Error killing process:", err)
		}
	case "ps":
		cmd := exec.Command("ps")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running ps command:", err)
		}
	default:
		cmd := exec.Command(command, commandArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
