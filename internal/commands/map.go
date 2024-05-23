package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nevnet99/pokedex-cli/internal/types"
)

var currentPage = 1
var locations = types.Response{}

func fetchLocations() error {
	res, resError := http.Get("https://pokeapi.co/api/v2/location-area" + fmt.Sprintf("?offset=%d", currentPage*20))

	if resError != nil {
		fmt.Println(resError)
		return resError
	}

	defer res.Body.Close()

	body, readError := io.ReadAll(res.Body)

	if readError != nil {
		fmt.Println(readError)
		return readError
	}

	response := types.Response{}

	err := json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err)
		return err
	}

	locations = response

	return nil
}

func MapFn() error {
	fmt.Printf("Current page: %d \n", currentPage)
	err := fetchLocations()

	if err != nil {
		return err
	}

	for i, location := range locations.Results {
		fmt.Printf("%d: %s \n", i+1, location.Name)
	}

	currentPage++

	return nil
}

func MapFnB() error {
	fmt.Printf("Current page: %d \n", currentPage)
	err := fetchLocations()

	if err != nil {
		return err
	}

	if currentPage == 1 {
		fmt.Println("You are already on the first page")
		return nil
	} else {
		currentPage--
	}

	for i, location := range locations.Results {
		fmt.Printf("%d: %s \n", i+1, location.Name)
	}

	return nil
}
