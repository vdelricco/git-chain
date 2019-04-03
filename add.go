package main

type add struct{}

func (command *add) help() string {
	return "Takes one arguments that is a git branch to add to the chain"
}

func (command *add) requiredArgs() int {
	return 1
}

func (command *add) execute(arguments []string) {
	addBranchToChain(arguments[0])
}
