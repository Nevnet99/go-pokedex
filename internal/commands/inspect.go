package commands

import (
	"errors"
	"fmt"

	"github.com/nevnet99/pokedex-cli/internal/types"
)

func Inspect(pokemon string, pokedex types.Pokedex) error {
	storedPokemon, ok := pokedex.Pokemon[pokemon]

	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v \n", storedPokemon.Name)
	fmt.Printf("Height: %v \n", storedPokemon.Height)
	fmt.Printf("Weight: %v \n", storedPokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range storedPokemon.Stats {
		fmt.Printf("-%v: %v \n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, pokemonType := range storedPokemon.Types {
		fmt.Printf("-%v \n", pokemonType.Type.Name)
	}

	return nil

}
