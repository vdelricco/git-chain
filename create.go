package main

type create struct{}

func (command *create) help() string {
	return "Takes two arguments, a base git branch and the first branch in the chain"
}

func (command *create) requiredArgs() int {
	return 2
}

func (command *create) execute(arguments []string) {
	createBranchChain(arguments)
}
