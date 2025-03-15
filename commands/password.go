package commands

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

type PasswordCommand struct{}

func (p *PasswordCommand) Execute(args []string) {
	length := 12
	special := false
	numbers := false
	uppercase := false

	if len(args) > 0 {
		if l, err := strconv.Atoi(args[0]); err == nil {
			length = l
		}
	}

	for _, arg := range args[1:] {
		switch arg {
		case "--special":
			special = true
		case "--numbers":
			numbers = true
		case "--uppercase":
			uppercase = true
		}
	}

	charset := "abcdefghijklmnopqrstuvwxyz"
	if uppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if numbers {
		charset += "0123456789"
	}
	if special {
		charset += "!@#$%^&*()_+-=[]{}|;:,.<>?"
	}

	password := make([]byte, length)
	for i := range password {
		randIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[randIndex.Int64()]
	}

	fmt.Println("Generated Password:", string(password))
}
