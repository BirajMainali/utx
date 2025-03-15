package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
)

type GuidCommand struct{}

func (g *GuidCommand) Execute(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: utx guid <version> [upper]")
		fmt.Println("Available versions: v1, v4, v5")
		fmt.Println("Use 'upper' as the second argument to get uppercase UUID.")
		return
	}

	version := args[0]
	toUpper := len(args) > 1 && strings.ToLower(args[1]) == "upper"

	var generatedUUID string

	switch version {
	case "v1":
		uuidV1, err := uuid.NewUUID()
		if err != nil {
			log.Fatalf("Error generating UUIDv1: %v", err)
		}
		generatedUUID = uuidV1.String()

	case "v4":
		generatedUUID = uuid.New().String()

	case "v5":
		if len(args) < 3 {
			fmt.Println("Usage: utx guid v5 <namespace> <name> [upper]")
			return
		}
		namespace := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(args[1]))
		uuidV5 := uuid.NewSHA1(namespace, []byte(args[2]))
		generatedUUID = uuidV5.String()

	default:
		fmt.Println("Invalid version. Available versions: v1, v4, v5")
		return
	}

	if toUpper {
		generatedUUID = strings.ToUpper(generatedUUID)
	}

	fmt.Println("Generated UUID:", generatedUUID)
}
