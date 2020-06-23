package model

// PokemonResponse is the struct used for the
// GET /pokemon/<pokemon> endpoint response
type PokemonResponse struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
