package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/nevnet99/pokedex-cli/internal/types"
)

var fetchedPokemon = types.PokemonResponse{}

func getPokemon(pokemon string) error {
	res, resError := http.Get("https://pokeapi.co/api/v2/pokemon/" + fmt.Sprintf(pokemon))

	if resError != nil {
		return resError
	}

	defer res.Body.Close()

	body, readError := io.ReadAll(res.Body)

	if readError != nil {
		return readError
	}

	response := types.PokemonResponse{}

	err := json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fetchedPokemon = response

	return nil
}

func getCatchResult(experience int) bool {
	catchPercentage := experience / 10

	if catchPercentage > 100 {
		catchPercentage = 100
	}

	randomNumber := rand.Intn(100) + 1

	fmt.Printf("Catch rate of %v%% | Random number: %v \n", catchPercentage, randomNumber)

	return randomNumber <= catchPercentage
}

func Catch(pokemon string, pokedex types.Pokedex) error {
	ok := getPokemon(pokemon)

	if ok != nil {
		return errors.New("failed to fetch pokemon")
	}

	fmt.Printf("Throwing a Pokeball at %v... \n", fetchedPokemon.Name)

	caught := getCatchResult(fetchedPokemon.BaseExperience)

	if caught {
		pokedex.Pokemon[pokemon] = fetchedPokemon

		fmt.Printf("%v was caught! \n", fetchedPokemon.Name)
	} else {
		fmt.Printf("%v escaped! \n", fetchedPokemon.Name)
	}

	return nil
}
