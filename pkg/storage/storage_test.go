package storage_test

import (
	"testing"

	"github.com/pcpratheesh/go-urlshortner/pkg/storage"
	"github.com/stretchr/testify/require"
)

func TestNewStorage(t *testing.T) {
	t.Run("initialization", func(t *testing.T) {
		_, err := storage.NewStorage(storage.STORE_TYPE_INMEMORY)
		require.Nil(t, err)
	})

	t.Run("not-implemented", func(t *testing.T) {
		_, err := storage.NewStorage("invalid-one")
		require.NotNil(t, err)
	})
}
