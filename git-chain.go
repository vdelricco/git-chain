package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Grab a pointer to the first input argument and check if it's nil
	if len(os.Args) < 2 {
		printCommandList()
		return
	}

	// The user provided a command argument; check if it matches something we can handle
	userDefinedCommand := getCommandFromUserInput(os.Args[1])
	if userDefinedCommand == nil {
		printCommandList()
		return
	}

	// Now that we have the command, grab the rest of the arguments the user provided
	commandArgs := os.Args[2:]

	// See if the user is requesting help
	if len(commandArgs) == 1 && commandArgs[0] == "help" {
		println("Help:\n")
		println(userDefinedCommand.help())
	}

	// See if the user provided the wrong amount of arguments
	if len(commandArgs) != userDefinedCommand.requiredArgs() {
		println("Wrong amount of arguments provided:\n")
		println(userDefinedCommand.help())
	}

	// Make sure we're inside of a git repo before we run any commands
	if err := exec.Command("git", "status").Run(); err != nil {
		println("Error! Ensure your current directory is inside a git repo")
		log.Fatal(err)
	}

	// Do the thing!
	userDefinedCommand.execute(commandArgs)
}
