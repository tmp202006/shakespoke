package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewShakespeareController(t *testing.T) {
	var assert = assert.New(t)
	cases := []struct {
		name     string
		input    string
		expected ShakespeareController
	}{
		{
			name:  "ShakespeareController is successfully created whith given URL",
			input: "FAKE_URL",
			expected: ShakespeareController{
				apiURL: "FAKE_URL",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(c.expected, NewShakespeareController(c.input))
		})
	}

}
