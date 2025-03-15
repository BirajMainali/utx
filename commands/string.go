package commands

import (
	"fmt"
	"strings"
)

type StringCommand struct{}

func (s *StringCommand) Execute(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: utx string <operation> <string1> [string2]")
		fmt.Println("Available operations: length, compare, upper, lower, trim, utc")
		return
	}

	operation := strings.ToLower(args[0])
	input := args[1]

	switch operation {
	case "length":
		fmt.Println("Length:", len(input))

	case "compare":
		if len(args) < 3 {
			fmt.Println("Usage: utx string compare <string1> <string2>")
			return
		}
		secondInput := args[2]
		if input == secondInput {
			fmt.Println("Strings are equal")
		} else {
			fmt.Println("Strings are not equal")
		}

	case "upper":
		fmt.Println("Uppercase:", strings.ToUpper(input))

	case "lower":
		fmt.Println("Lowercase:", strings.ToLower(input))

	case "trim":
		fmt.Println("Trimmed:", strings.TrimSpace(input))

	default:
		fmt.Println("Invalid operation. Available operations: length, compare, upper, lower, trim, utc")
	}
}
