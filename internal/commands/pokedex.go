package commands

import (
	"errors"
	"fmt"

	"github.com/nevnet99/pokedex-cli/internal/types"
)

func PokedexFn(pokedex types.Pokedex) error {

	if len(pokedex.Pokemon) == 0 {
		return errors.New("go catch some pokemon with the `catch` command")
	}

	for _, pokemon := range pokedex.Pokemon {
		fmt.Printf("-%v \n", pokemon.Name)
	}

	return nil
}
