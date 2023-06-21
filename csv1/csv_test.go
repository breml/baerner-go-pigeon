package csv

import (
	"testing"

	"github.com/matryer/is"
)

func TestCSV(t *testing.T) {
	tt := []struct {
		name  string
		input string

		wantErr bool
	}{
		{
			name:  "one line",
			input: `a,b,c`,

			wantErr: false,
		},
		{
			name: "multiple lines",
			input: `a,b,c
d,e,g
g,h,i`,

			wantErr: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)

			ret, err := Parse("", []byte(tc.input))
			is.True((err != nil) == tc.wantErr)

			lines := toAnySlice(ret)
			for l, line := range lines {
				vals := toAnySlice(line)
				for v, val := range vals {
					s := val.(string)
					t.Log(l, v, s)
				}
			}
		})
	}
}
