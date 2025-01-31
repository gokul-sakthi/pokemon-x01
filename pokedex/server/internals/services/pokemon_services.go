package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

// FetchPokedex retrieves Pokémon data from the PokeAPI.
func FetchPokedex() (interface{}, error) {
	resource, err := pokeapi.Resource("pokedex")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch pokedex: %w", err)
	}

	return &resource.Results, nil
}

type FetchPokemonOpts struct {
	Offset int
	Limit  int
}

// Get paginated data from the poke api
func FetchPokemon(opts *FetchPokemonOpts) (interface{}, error) {
	if opts == nil {
		opts = &FetchPokemonOpts{Offset: 0, Limit: 10}
	}

	resource, err := pokeapi.Resource("pokemon", opts.Offset, opts.Limit)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch pokemon: %w", err)
	}

	return &resource.Results, nil
}

func FetchPokemonByID(id int) (interface{}, error) {
	resource, err := pokeapi.Pokemon(strconv.Itoa(id))

	if err != nil {
		return nil, fmt.Errorf("failed to fetch pokemon: %w", err)
	}

	return &resource, nil
}

func SearchPokemon(identifier string) (interface{}, error) {

	const startsWith = "^"
	resource, err := pokeapi.Search("pokemon", startsWith+strings.ToLower(identifier))

	if err != nil {
		return nil, fmt.Errorf("failed to search pokemon: %w", err)
	}

	return &resource, nil
}
