package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var AliasFile = filepath.Join(os.Getenv("HOME"), ".aliases.json")

type AliasMap map[string]string

type AliasCommand struct{}

func (a *AliasCommand) IsAlias(name string) bool {
	aliases, _ := loadAliases()
	_, exists := aliases[name]
	return exists
}

func ensureAliasFile() {
	if _, err := os.Stat(AliasFile); os.IsNotExist(err) {
		os.WriteFile(AliasFile, []byte("{}"), 0644)
	}
}

func loadAliases() (AliasMap, error) {
	ensureAliasFile()
	file, err := os.ReadFile(AliasFile)
	if err != nil {
		return nil, err
	}

	var aliases AliasMap
	err = json.Unmarshal(file, &aliases)
	if err != nil {
		return nil, err
	}
	return aliases, nil
}

func saveAliases(aliases AliasMap) error {
	data, err := json.MarshalIndent(aliases, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(AliasFile, data, 0644)
}

func (a *AliasCommand) Execute(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: alias <add|remove|list|run> [name] [command]")
		return
	}

	operation := args[0]

	switch operation {
	case "add":
		if len(args) < 3 {
			fmt.Println("Usage: alias add <name> <command>")
			return
		}
		name, command := args[1], strings.Join(args[2:], " ")
		addAlias(name, command)
	case "remove":
		if len(args) < 2 {
			fmt.Println("Usage: alias remove <name>")
			return
		}
		removeAlias(args[1])
	case "list":
		listAliases()
	default:
		executeAlias(operation)
	}
}

func addAlias(name, command string) {
	aliases, _ := loadAliases()
	aliases[name] = command
	saveAliases(aliases)
	fmt.Printf("Alias '%s' added for command: %s\n", name, command)
}

func listAliases() {
	aliases, _ := loadAliases()
	fmt.Println("Registered Aliases:")
	for name, command := range aliases {
		fmt.Printf("%s -> %s\n", name, command)
	}
}

func removeAlias(name string) {
	aliases, _ := loadAliases()
	if _, exists := aliases[name]; exists {
		delete(aliases, name)
		saveAliases(aliases)
		fmt.Printf("Alias '%s' removed\n", name)
	} else {
		fmt.Printf("Alias '%s' not found\n", name)
	}
}

func executeAlias(name string) {
	aliases, _ := loadAliases()
	if command, exists := aliases[name]; exists {
		runCommand(command)
	} else {
		fmt.Printf("Alias '%s' not found\n", name)
	}
}

func runCommand(command string) {
	parts := strings.Split(command, " ")
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}
