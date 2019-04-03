package main

type command interface {
	help() string
	requiredArgs() int
	execute(arguments []string)
}

var commandStrings = []string{"add", "create", "list", "sync"}

func printCommandList() {
	println("Valid commands:")
	for _, command := range commandStrings {
		println(command)
	}
}

func getCommandFromUserInput(input string) command {
	switch input {
	case "add":
		return &add{}
	case "create":
		return &create{}
	case "list":
		return &list{}
	case "sync":
		return &sync{}
	default:
		return nil
	}
}
