package main

import (
	"fmt"
	"os"

	"github.com/birajmainali/utx/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: utx <command> <args>")
		return
	}

	commandName := os.Args[1]
	args := os.Args[2:]

	// Load aliases and check if the command is an alias
	aliasCmd := &commands.AliasCommand{}
	if aliasCmd.IsAlias(commandName) {
		aliasCmd.Execute([]string{commandName})
		return
	}

	invoker := NewInvoker()
	invoker.RegisterCommand("guid", &commands.GuidCommand{})
	invoker.RegisterCommand("str", &commands.StringCommand{})
	invoker.RegisterCommand("date", &commands.DateCommand{})
	invoker.RegisterCommand("json", &commands.JsonCommand{})
	invoker.RegisterCommand("net", &commands.NetworkCommand{})
	invoker.RegisterCommand("pwd", &commands.PasswordCommand{})
	invoker.RegisterCommand("help", &commands.HelpCommand{})
	invoker.RegisterCommand("alias", aliasCmd)

	invoker.ExecuteCommand(commandName, args)
}
