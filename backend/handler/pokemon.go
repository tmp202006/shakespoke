package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tmp202006/shakespoke/backend/model"
)

// GetPokemon is the handler for the GET /pokemon/<pokemon> endpoint
// For the given pokemon, returns its Shakesperian description
func (h Handler) GetPokemon(w http.ResponseWriter, r *http.Request) {
	pokemonName := chi.URLParam(r, "pokemon")

	if pokemonName == "" {
		errorResponse := model.ErrorResponse{
			Code:    model.ErrMissingInput,
			Message: model.ErrMessageMissingInput,
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			errorResponse,
		)
	}

	// Retrieve Pokemon
	pokemon, err := h.PokemonController.GetPokemon(pokemonName)
	if err != nil {
		errorResponse := model.ErrorResponse{
			Error:   err,
			Code:    model.ErrGetPokemon,
			Message: model.ErrMessageGetPokemon,
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			errorResponse,
		)
		return
	}

	// Retrieve Pokemon Species
	species, err := h.PokemonController.GetPokemonSpecies(pokemon.Species.URL)
	if err != nil {
		errorResponse := model.ErrorResponse{
			Error:   err,
			Code:    model.ErrGetPokemonSpecies,
			Message: model.ErrMessageGetPokemonSpecies,
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			errorResponse,
		)
		return
	}

	if len(species.FlavorTextExtries) == 0 {
		errorResponse := model.ErrorResponse{
			Code:    model.ErrNoDescriptions,
			Message: model.ErrMessageNoDescriptions,
		}
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(
			errorResponse,
		)
		return
	}

	// Translate the description
	// For the sake of this project, pick the first available item
	flavorText := h.escapeText(species.FlavorTextExtries[0].FlavorText)
	translation, err := h.ShakespeareController.Translate(flavorText)
	if err != nil {
		errorResponse := model.ErrorResponse{
			Error:   err,
			Code:    model.ErrGetTranslation,
			Message: model.ErrMessageGetTranslation,
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			errorResponse,
		)
		return
	}

	if translation.Success.Total == 0 {
		errorResponse := model.ErrorResponse{
			Error:   err,
			Code:    model.ErrNoTranslation,
			Message: model.ErrMessageNoTranslation,
		}
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(
			errorResponse,
		)
		return
	}

	// Return the response
	response := model.PokemonResponse{
		Name:        pokemonName,
		Description: &translation.Contents.Translated,
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(
		response,
	)
}
