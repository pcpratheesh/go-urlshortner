package storage

import "fmt"

type inMemory struct{}

var memoryStorageMap = make(map[string]string, 0)

func NewInMemoryStorage() StorageInterface {
	return &inMemory{}
}

// save the url in memory
func (in *inMemory) SaveURL(shortURL, originalURL string) error {
	if _, ok := memoryStorageMap[shortURL]; ok {
		return fmt.Errorf("shorten url already exists in the memory")
	}

	// else save the data
	memoryStorageMap[shortURL] = originalURL

	return nil
}

// retrieve the url from memory
func (in *inMemory) RetrieveURL(shortURL string) string {
	return memoryStorageMap[shortURL]
}

// retrieve the url from memory
func (in *inMemory) CheckURLExists(originalURL string) (string, bool) {
	for short, orign := range memoryStorageMap {
		if orign == originalURL {
			return short, true
		}
	}
	return "", false
}
