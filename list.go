package main

type list struct{}

func (command *list) help() string {
	return "Takes no arguments, lists the branches in the chain"
}

func (command *list) requiredArgs() int {
	return 0
}

func (command *list) execute(_ []string) {
	for _, branch := range getBranchChain() {
		println(branch)
	}
}
