package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPokemonController(t *testing.T) {
	var assert = assert.New(t)
	cases := []struct {
		name     string
		input    string
		expected PokemonController
	}{
		{
			name:  "PokemonController is successfully created whith given URL",
			input: "FAKE_URL",
			expected: PokemonController{
				apiURL: "FAKE_URL",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(c.expected, NewPokemonController(c.input))
		})
	}

}
