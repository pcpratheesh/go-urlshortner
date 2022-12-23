package shortner

import (
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

var (
	RandomStringLength = 15
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + strings.Replace(uuid.New().String(), "-", "", -1))

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func GenerateShortLink(url string) string {
	return string(RandomString(RandomStringLength))
}
