package controller

import "github.com/tmp202006/shakespoke/backend/model"

// PokemonControllerMock is the mock for PokemonControl
type PokemonControllerMock struct {
	Pokemon             *model.Pokemon
	PokemonError        error
	PokemonSpecies      *model.PokemonSpecies
	PokemonSpeciesError error
}

//GetPokemon mocks the GetPokemon function
func (mock PokemonControllerMock) GetPokemon(pokemonName string) (*model.Pokemon, error) {
	return mock.Pokemon, mock.PokemonError
}

//GetPokemonSpecies mocks the GetPokemonSpecies function
func (mock PokemonControllerMock) GetPokemonSpecies(speciesURL string) (*model.PokemonSpecies, error) {
	return mock.PokemonSpecies, mock.PokemonSpeciesError
}
