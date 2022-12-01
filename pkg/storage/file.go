package storage

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type fileStorage struct {
	Filepath string
}

func NewFileStorage(filepath string) StorageInterface {
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return &fileStorage{filepath}
}

// save the url in memory
func (store *fileStorage) SaveURL(shortURL, originalURL string) error {
	// check the string already exists
	b, err := ioutil.ReadFile(store.Filepath) // just pass the file name
	if err != nil {
		return err
	}

	if strings.Contains(string(b), shortURL) {
		return fmt.Errorf("shorten url already exists in the memory")
	}

	err = ioutil.WriteFile(store.Filepath, []byte(fmt.Sprintf("%s = %s;\n", shortURL, originalURL)), 0644)

	return err
}

// retrieve the url from memory
func (store *fileStorage) RetrieveURL(shortURL string) string {
	f, err := os.Open(store.Filepath)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(customSplitFunc)

	for scanner.Scan() {
		storageData := scanner.Text()

		splitedSlice := strings.Split(storageData, " = ")

		if splitedSlice[0] == shortURL {
			return splitedSlice[1]
		}
	}

	return ""
}

func customSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if atEOF {
		return len(data), data, nil
	}

	if i := strings.Index(string(data), ";\n"); i >= 0 {
		//skip the delimiter in advancing to the next pair
		return i + 2, data[0:i], nil
	}
	return
}

// retrieve the url from memory
func (store *fileStorage) CheckURLExists(originalURL string) (string, bool) {
	f, err := os.Open(store.Filepath)
	if err != nil {
		logrus.Error(err)
		return "", false
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(customSplitFunc)

	for scanner.Scan() {
		storageData := scanner.Text()

		splitedSlice := strings.Split(storageData, " = ")

		if splitedSlice[1] == originalURL {
			return splitedSlice[0], true
		}
	}

	return "", false
}
