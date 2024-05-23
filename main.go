package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nevnet99/pokedex-cli/internal/commands"
	"github.com/nevnet99/pokedex-cli/internal/types"
)

func main() {
	cliCommands := make(map[string]types.CliCommand)

	cliCommands["help"] = types.CliCommand{
		Name:        "help",
		Description: "Prints the list of available commands",
		Callback: func() error {
			return commands.Help(cliCommands)
		},
	}

	cliCommands["exit"] = types.CliCommand{
		Name:        "exit",
		Description: "Exits the Poxedex.",
		Callback:    commands.Exit,
	}

	cliCommands["map"] = types.CliCommand{
		Name:        "map",
		Description: "Prints the map of the region.",
		Callback:    commands.MapFn,
	}

	cliCommands["mapb"] = types.CliCommand{
		Name:        "mapb",
		Description: "Prints the map of the region.",
		Callback:    commands.MapFnB,
	}

	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			fmt.Print("Pokedex > ")
			continue
		}

		if command, ok := cliCommands[text]; ok {
			err := command.Callback()

			if err != nil {
				if err.Error() == "exit" {
					break
				}
				fmt.Println("Error:", err)
			}

			fmt.Print("Pokedex > ")

		} else {
			fmt.Printf("Unknown command: %s please run the command 'help'", text)
			continue
		}
	}

}
