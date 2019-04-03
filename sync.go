package main

type sync struct{}

func (command *sync) help() string {
	return "Takes no arguments, syncs the branches from the base of the chain"
}

func (command *sync) requiredArgs() int {
	return 0
}

func (command *sync) execute(_ []string) {
	branches := getBranchChain()

	runCommand("git checkout " + branches[0])

	runCommand("git rebase " + branches[0] + " " + branches[1])

	for _, branch := range branches[2:] {
		println("syncing:", branch)
		runCommand("git rebase --onto HEAD ORIG_HEAD " + branch)
	}

	runCommand("git checkout " + branches[0])
}
