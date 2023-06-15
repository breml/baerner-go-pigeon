package aparser

import (
	"testing"

	"github.com/matryer/is"
)

func TestWordlist(t *testing.T) {
	tt := []struct {
		name  string
		input string

		wantErr bool
	}{
		{
			name:  "one a",
			input: "a",

			wantErr: false,
		},
		{
			name:  "two a",
			input: "aa",

			wantErr: false,
		},
		{
			name:  "many a",
			input: "aaaaaaaaaaaaa",

			wantErr: false,
		},
		{
			name:  "not an a",
			input: "b",

			wantErr: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)

			_, err := Parse("", []byte(tc.input))
			is.True((err != nil) == tc.wantErr)
		})
	}
}
