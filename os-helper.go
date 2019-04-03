package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func createFile(filePath string) *os.File {
	file, err := os.Create(gitDataFile)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func openFile(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil
	}

	return file
}

func runCommand(command string) {
	commandArgs := strings.Fields(command)

	out, err := exec.Command(commandArgs[0], commandArgs[1:]...).CombinedOutput()
	if err != nil {
		log.Fatal(string(out))
	}
}
