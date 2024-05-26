package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/nevnet99/pokedex-cli/internal/commands"
	"github.com/nevnet99/pokedex-cli/internal/types"
)

func main() {
	cliCommands := make(map[string]types.CliCommand)

	cliCommands["help"] = types.CliCommand{
		Name:        "help",
		Description: "Prints the list of available commands",
		Callback: func(context []string) error {
			return commands.Help(cliCommands)
		},
	}

	cliCommands["exit"] = types.CliCommand{
		Name:        "exit",
		Description: "Exits the Poxedex.",
		Callback: func(context []string) error {
			return commands.Exit()
		},
	}

	cliCommands["map"] = types.CliCommand{
		Name:        "map",
		Description: "Prints the map of the region.",
		Callback: func(context []string) error {
			return commands.MapFn()
		},
	}

	cliCommands["mapb"] = types.CliCommand{
		Name:        "mapb",
		Description: "Prints the map of the region.",
		Callback: func(context []string) error {
			return commands.MapFnB()
		},
	}

	cliCommands["explore"] = types.CliCommand{
		Name:        "explore",
		Description: "Traverse areas and find out what pokemon can be found on that location",
		Callback: func(context []string) error {
			if len(context) == 1 {
				return errors.New("location parameter is required")
			}

			location := context[1]

			return commands.Explore(location)
		},
	}

	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()
		parameters := strings.Split(text, " ")
		commandRan := parameters[0]

		if text == "" {
			fmt.Print("Pokedex > ")
			continue
		}

		if command, ok := cliCommands[commandRan]; ok {

			err := command.Callback(parameters)

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
