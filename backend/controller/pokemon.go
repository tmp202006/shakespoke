package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tmp202006/shakespoke/backend/model"
)

// PokemonControl is the interface defining the methods for interacting
// with Pokemon external API
type PokemonControl interface {
	GetPokemon(pokemonName string) (*model.Pokemon, error)
	GetPokemonSpecies(speciesURL string) (*model.PokemonSpecies, error)
}

// PokemonController implements PokemonControl by actually
// connecting to the external API
type PokemonController struct {
	apiURL string
}

// NewPokemonController a new PokemonController instance
func NewPokemonController(apiURL string) PokemonController {
	return PokemonController{
		apiURL: apiURL,
	}
}

// GetPokemon implements the request to the external API for retrieving
// the details about a pokemon
func (pc PokemonController) GetPokemon(pokemonName string) (*model.Pokemon, error) {
	url := fmt.Sprintf("%v/pokemon/%v", pc.apiURL, pokemonName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var pokemon model.Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return nil, err
	}
	return &pokemon, nil
}

// GetPokemonSpecies implements the request to the external API for
// retrieving the details about a pokemon species
func (pc PokemonController) GetPokemonSpecies(speciesURL string) (*model.PokemonSpecies, error) {
	resp, err := http.Get(speciesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var species model.PokemonSpecies
	if err := json.Unmarshal(body, &species); err != nil {
		return nil, err
	}
	return &species, nil
}
