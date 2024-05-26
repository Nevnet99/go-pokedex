package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/nevnet99/pokedex-cli/internal/pokecache"
	"github.com/nevnet99/pokedex-cli/internal/types"
)

var cache2 = pokecache.NewCache(time.Second * 5)
var pokemonInArea = types.PokemonAreaResponse{}

func fetchPokemonAtLocation(location string) error {
	cacheKey := fmt.Sprintf("%v", currentPage)

	if cached, found := cache2.Get(cacheKey); found {
		fmt.Println("*>*>*> USED THE CACHE <*<*<*")

		response := types.PokemonAreaResponse{}

		err := json.Unmarshal(cached, &response)

		if err != nil {
			fmt.Println(err)
			return err
		}

		pokemonInArea = response

		return nil
	}

	res, resError := http.Get("https://pokeapi.co/api/v2/location-area/" + fmt.Sprintf(location) + fmt.Sprintf("?offset=%d", currentPage*20))

	if resError != nil {
		fmt.Println(resError)
		return resError
	}

	defer res.Body.Close()

	body, readError := io.ReadAll(res.Body)

	cache2.Add(fmt.Sprintf("%v", currentPage), body)

	if readError != nil {
		return readError
	}

	response := types.PokemonAreaResponse{}

	err := json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err)
		return err
	}

	pokemonInArea = response

	return nil
}

func Explore(location string) error {
	fmt.Printf("Exploring %v...\n", location)
	err := fetchPokemonAtLocation(location)

	if err != nil {
		return err
	}

	fmt.Print("Found Pokemon:\n")
	for _, encounter := range pokemonInArea.PokemonEncounters {
		fmt.Printf("- %v\n", encounter.Pokemon.Name)
	}
	return nil

}
