package commands

import (
	"fmt"

	"github.com/nevnet99/pokedex-cli/internal/types"
)

func Help(commands map[string]types.CliCommand) error {
	fmt.Println("Available commands:")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}
