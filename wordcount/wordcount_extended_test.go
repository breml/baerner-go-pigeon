package wordcount_test

import (
	"testing"

	"github.com/matryer/is"

	"github.com/breml/baerner-go-pigeon/wordcount"
)

func TestWordcountExtended(t *testing.T) {
	t.Skip("skipped")

	tt := []struct {
		name  string
		words int
	}{
		{
			name:  "Welcome Bärner Go audience",
			words: 4,
		},
		{
			name: `Lorem ipsum dolor sit
amet consetetur sadipscing elitr
sed diam nonumy eirmod
tempor invidunt ut labore
et dolore magna aliquyam
erat sed diam voluptua`,
			words: 24,
		},
		{
			name:  "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.",
			words: 24,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)

			res, err := wordcount.Parse("", []byte(tc.name))
			is.NoErr(err)

			wordCount, ok := res.(int)
			is.True(ok)
			is.Equal(tc.words, wordCount)
		})
	}
}
