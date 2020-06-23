package model

const (
	// ErrMissingInput is the error for missing input
	ErrMissingInput = "missing-input"
	// ErrMessageMissingInput is the error message for missing input
	ErrMessageMissingInput = "The provided input is missing, empty, or undefined"
	// ErrGetPokemon is the error for get pokemon
	ErrGetPokemon = "error-get-pokemon"
	// ErrMessageGetPokemon is the error message for get pokemon
	ErrMessageGetPokemon = "Failed to request the GET /pokemon/<pokemon> API"
	// ErrGetPokemonSpecies is the error for get pokemon species
	ErrGetPokemonSpecies = "error-get-pokemon-species"
	// ErrMessageGetPokemonSpecies is the error message for get pokemon species
	ErrMessageGetPokemonSpecies = "Failed to request the GET /pokemon-species/<species> API"
	// ErrGetTranslation is the error for get translation
	ErrGetTranslation = "error-get-translation"
	// ErrMessageGetTranslation is the error message for get translation
	ErrMessageGetTranslation = "Failed to request the POST /translation API"
	// ErrNoDescriptions is the error returned when no description (flavor text) is available
	ErrNoDescriptions = "error-no-descriptions"
	// ErrMessageNoDescriptions is the error message returned when no description (flavor text) is available
	ErrMessageNoDescriptions = "No description (flavor text) available for the given pokemon"
	// ErrNoTranslation is the error returned when no translation si available
	ErrNoTranslation = "error-no-translation"
	// ErrMessageNoTranslation is the error messagereturned when no translation si available
	ErrMessageNoTranslation = "No translation available for the given pokemon"
)

// ErrorResponse represents a generic error response
type ErrorResponse struct {
	Error   error  `json:"error"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
