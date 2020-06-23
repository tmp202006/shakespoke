package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/tmp202006/shakespoke/backend/controller"
	"github.com/tmp202006/shakespoke/backend/handler"
)

func main() {
	// Init environment variables
	pokemonAPIURL := os.Getenv("POKEMON_API_URL")
	shakespeareAPIURL := os.Getenv("SHAKESPEARE_API_URL")
	portNumber, err := strconv.ParseInt(os.Getenv("PORT_NUMBER"), 10, 64)
	if err != nil {
		log.Fatalf("%v: error parsing PORT_NUMBER", err)
		return
	}
	port := fmt.Sprintf(":%v", portNumber)

	// Init controllers and handler
	pc := controller.NewPokemonController(pokemonAPIURL)
	sc := controller.NewShakespeareController(shakespeareAPIURL)
	h := handler.Handler{
		PokemonController:     &pc,
		ShakespeareController: &sc,
	}

	log.Printf("[INFO] Pokemon API URL: %v", pokemonAPIURL)
	log.Printf("[INFO] Shakespeare API URL: %v", shakespeareAPIURL)
	log.Printf("[INFO] Port Number: %v", portNumber)
	log.Print("[INFO] Starting router")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}))
	r.Get("/pokemon/{pokemon}", h.GetPokemon)
	http.ListenAndServe(port, r)
}
