package storage_test

import (
	"testing"

	"github.com/pcpratheesh/go-urlshortner/pkg/storage"
	"github.com/stretchr/testify/require"
)

func TestMemoryStorage(t *testing.T) {
	t.Run("init-storage-object", func(t *testing.T) {
		store, err := storage.NewStorage(storage.STORE_TYPE_INMEMORY)
		require.Nil(t, err)

		t.Run("store-url", func(t *testing.T) {
			err := store.SaveURL("generated-url", "original-url.com")
			require.Nil(t, err)
		})

		t.Run("store-existing-url", func(t *testing.T) {
			err := store.SaveURL("generated-url", "original-url.com")
			require.NotNil(t, err)
		})

		t.Run("retrieve-url", func(t *testing.T) {
			originalURL := store.RetrieveURL("generated-url")
			require.NotEmpty(t, originalURL)

			require.Equal(t, originalURL, "original-url.com")
		})

		t.Run("retrieve-not-existing-url", func(t *testing.T) {
			originalURL := store.RetrieveURL("generated-url-not-existing")
			require.Empty(t, originalURL)
		})

		t.Run("check-url-exists", func(t *testing.T) {
			_, ok := store.CheckURLExists("original-url.com")
			require.NotNil(t, ok)
		})

		t.Run("check-url-not-exists", func(t *testing.T) {
			_, ok := store.CheckURLExists("original-url-not-exists.com")
			require.False(t, ok)
		})
	})
}
