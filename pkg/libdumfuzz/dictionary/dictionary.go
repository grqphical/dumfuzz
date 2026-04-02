package dictionary

import (
	_ "embed"
	"math/rand/v2"
	"strings"
)

//go:embed dictionary.txt
var dictionary []byte

var Words []string

func GetRandomWord() string {
	idx := rand.IntN(len(Words))
	return Words[idx]
}

func GetRandomWordOfLength(length int) string {
	word := ""
	for len(word) != length {
		word = GetRandomWord()
	}

	return word
}

func init() {
	lines := strings.SplitSeq(string(dictionary), "\r\n")

	for line := range lines {
		if len(line) == 0 {
			continue
		}
		Words = append(Words, line)
	}
}
