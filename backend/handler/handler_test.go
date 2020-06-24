package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscapeText(t *testing.T) {
	var (
		assert = assert.New(t)
		h      = Handler{}
	)

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string returns an empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "random sentence returns the same sentence",
			input:    "lorem ipsum dolor sit amet",
			expected: "lorem ipsum dolor sit amet",
		},
		{
			name:     "carriage returns and new lines are replaced with empy spaces",
			input:    "lorem \nipsum \r\ndolor \nsit \r\namet",
			expected: "lorem  ipsum  dolor  sit  amet",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(c.expected, h.escapeText(c.input))
		})
	}
}
