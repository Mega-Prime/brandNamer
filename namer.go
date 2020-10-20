package brandNamer

import (
	"math/rand"
	"time"
)

const (
	consonants = "bcdfghjklmnpqrstvwxyz"
	vowels     = "aeiou"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func Name(length int) string {
	var res string

	for i := 0; i < length; i++ {
		res += string(consonants[seededRand.Intn(len(consonants))])
		res += string(vowels[seededRand.Intn(len(vowels))])
	}
	return res
}
