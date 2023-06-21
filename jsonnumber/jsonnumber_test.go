package jsonnumber

import (
	"testing"

	"github.com/matryer/is"
)

func TestJSONNumber(t *testing.T) {
	tt := []struct {
		name  string
		input string

		wantErr bool
	}{
		{
			name:  "int",
			input: "3",

			wantErr: false,
		},
		{
			name:  "float",
			input: "3.1415",

			wantErr: false,
		},
		{
			name:  "float with exponent",
			input: "3.1415e1",

			wantErr: false,
		},
		{
			name:  "not a json number",
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
