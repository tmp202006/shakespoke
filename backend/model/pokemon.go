package model

// Pokemon represents a pokemon resource as defined in https://pokeapi.co/
// For the sake of this project, only part of the fields have been replicated
type Pokemon struct {
	Name    string                  `json:"name"`
	Species PokemonSpeciesReference `json:"species"`
}

// PokemonSpeciesReference references the resource for the pokemon species
type PokemonSpeciesReference struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonSpecies represents a pokemon species
// For the sake of this project, only part of the fields have been replicated
type PokemonSpecies struct {
	FlavorTextExtries []FlavorTextEntry `json:"flavor_text_entries"`
}

// FlavorTextEntry a single entry representing a flavor text
// For the sake of this project, only part of the fields have been replicated
type FlavorTextEntry struct {
	FlavorText string `json:"flavor_text"`
}
