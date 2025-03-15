package main

import "fmt"

type CommandInvoker struct {
	commands map[string]ICommand
}

func NewInvoker() *CommandInvoker {
	return &CommandInvoker{commands: make(map[string]ICommand)}
}

func (ci *CommandInvoker) RegisterCommand(name string, command ICommand) {
	ci.commands[name] = command
}

func (ci *CommandInvoker) ExecuteCommand(name string, args []string) {
	if command, exists := ci.commands[name]; exists {
		command.Execute(args)
	} else {
		fmt.Println("Unknown command:", name)
	}
}
