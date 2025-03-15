package commands

import (
	"encoding/json"
	"fmt"
)

type JsonCommand struct{}

func (j *JsonCommand) Execute(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: utx json <operation> <json_string>")
		fmt.Println("Available operations: validate, format")
		return
	}

	operation := args[0]
	jsonStr := args[1]

	switch operation {

	case "validate":
		var js map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &js); err != nil {
			fmt.Println("Invalid JSON:", err)
			return
		}
		fmt.Println("Valid JSON ✅")

	case "format":
		var js map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &js); err != nil {
			fmt.Println("Invalid JSON:", err)
			return
		}
		formattedJSON, err := json.MarshalIndent(js, "", "  ")
		if err != nil {
			fmt.Println("Error formatting JSON:", err)
			return
		}
		fmt.Println(string(formattedJSON))

	default:
		fmt.Println("Invalid operation. Available operations: validate, format")
	}
}
