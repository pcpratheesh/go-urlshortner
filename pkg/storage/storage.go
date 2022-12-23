package storage

import (
	"fmt"
)

type StorageInterface interface {
	CheckURLExists(string) (string, bool)
	SaveURL(string, string) error
	RetrieveURL(string) string
}

var (
	STORE_TYPE_INMEMORY = "InMemory"
	STORE_TYPE_FILE     = "File"
)

// initiate new storage
func NewStorage(store string) (StorageInterface, error) {
	switch store {
	case STORE_TYPE_INMEMORY:
		return NewInMemoryStorage(), nil
	case STORE_TYPE_FILE:
		return NewFileStorage("data/contents.txt"), nil
	default:
		return nil, fmt.Errorf("%v store is not implemented", store)
	}

}
