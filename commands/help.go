package commands

import "fmt"

type HelpCommand struct{}

func (h *HelpCommand) Execute(args []string) {
	helpText := `
Usage: utx <command> [options]

Available Commands:
  date          Perform date operations (now, utc, add, diff)
  guid          Generate GUIDs (v1, v4, v5) with optional --upper flag
  json          Validate or format JSON
  net           Network utilities (local-ip, public-ip, ping, dns, interfaces)
  pwd           Generate a secure password <length> [--special --numbers --uppercase]
  str           String operations (length, compare, upper, lower, trim)
  help          Show this help message
`
	fmt.Println(helpText)
}
