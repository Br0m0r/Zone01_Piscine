package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	// Attempt to read from stdin for piped input
	stdinBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read from stdin: %v\n", err)
		os.Exit(1)
	}
	inputStr := string(stdinBytes)

	cmd := exec.Command("./quadchecker", inputStr)

	// Create a pipe for the command's standard output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe: ", err.Error())
		return
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command: ", err.Error())
		return
	}

	// Read the output from the command
	output, err := io.ReadAll(stdout)
	if err != nil {
		fmt.Println("Error reading command output: ", err.Error())
		return
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		fmt.Println("Command finished with error: ", err.Error())
		return
	}

	fmt.Print(string(output))
}
