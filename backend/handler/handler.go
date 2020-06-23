package handler

import (
	"html"
	"regexp"

	c "github.com/tmp202006/shakespoke/backend/controller"
)

// Handler serves and responds to API requests
type Handler struct {
	PokemonController     *c.PokemonController
	ShakespeareController *c.ShakespeareController
}

// escapeText makes the text url safe
func (h Handler) escapeText(s string) string {
	re := regexp.MustCompile(`\r?\n`)
	s = html.UnescapeString(s)
	return re.ReplaceAllString(s, " ")
}
