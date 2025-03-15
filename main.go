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

	invoker := NewInvoker()
	invoker.RegisterCommand("guid", &commands.GuidCommand{})
	invoker.RegisterCommand("str", &commands.StringCommand{})
	invoker.RegisterCommand("date", &commands.DateCommand{})
	invoker.RegisterCommand("json", &commands.JsonCommand{})
	invoker.RegisterCommand("net", &commands.NetworkCommand{})
	invoker.RegisterCommand("pwd", &commands.PasswordCommand{})

	invoker.ExecuteCommand(os.Args[1], os.Args[2:])
}
