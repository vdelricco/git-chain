package main

import (
	"bufio"
	"log"
)

var gitDataFile = "/tmp/git-chain-data"

func createBranchChain(branches []string) {
	file := createFile(gitDataFile)
	defer file.Close()

	for _, branch := range branches {
		file.WriteString(branch + "\n")
	}

	file.Sync()
}

func addBranchToChain(branch string) {
	file := openFile(gitDataFile)

	if file == nil {
		println("No git chain exists. Use the create command to get one started.")
		return
	}
	defer file.Close()

	println("Writing branch:", branch)
	file.WriteString(branch + "\n")

	file.Sync()
}

func getBranchChain() []string {
	file := openFile(gitDataFile)

	if file == nil {
		println("No git chain exists. Use the create command to get one started.")
		return []string{}
	}
	defer file.Close()

	var branchChain []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		branchChain = append(branchChain, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return branchChain
}
