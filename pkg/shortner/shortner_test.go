package shortner_test

import (
	"testing"

	"github.com/pcpratheesh/go-urlshortner/pkg/shortner"
	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	t.Run("generate-random-string", func(t *testing.T) {
		str := shortner.RandomString(10)
		require.NotEmpty(t, str)

		require.Equal(t, len(str), 10)
	})

	t.Run("generate-short-link", func(t *testing.T) {
		str := shortner.GenerateShortLink("abc")
		require.NotEmpty(t, str)

		require.Equal(t, len(str), shortner.RandomStringLength)
	})
}
