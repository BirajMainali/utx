package commands

import (
	"fmt"
	"strconv"
	"time"
)

type DateCommand struct{}

func (d *DateCommand) Execute(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: utx date <operation> [arguments]")
		fmt.Println("Available operations: now, utc, add, diff")
		return
	}

	operation := args[0]

	switch operation {
	case "now":
		fmt.Println("Current Date/Time:", time.Now().Format("2006-01-02 15:04:05"))

	case "utc":
		fmt.Println("Current UTC time:", time.Now().UTC())

	case "add":
		if len(args) < 3 {
			fmt.Println("Usage: utx date add <YYYY-MM-DD> <days>")
			return
		}
		date, err := time.Parse("2006-01-02", args[1])
		if err != nil {
			fmt.Println("Invalid date format. Use YYYY-MM-DD.")
			return
		}
		days, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid number of days.")
			return
		}
		newDate := date.AddDate(0, 0, days)
		fmt.Println("New Date:", newDate.Format("2006-01-02"))

	case "diff":
		if len(args) < 3 {
			fmt.Println("Usage: utx date diff <YYYY-MM-DD> <YYYY-MM-DD>")
			return
		}
		date1, err1 := time.Parse("2006-01-02", args[1])
		date2, err2 := time.Parse("2006-01-02", args[2])
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid date format. Use YYYY-MM-DD.")
			return
		}
		diff := date2.Sub(date1).Hours() / 24
		fmt.Printf("Difference: %.0f days\n", diff)

	default:
		fmt.Println("Invalid operation. Available operations: now, format, add, diff")
	}
}
