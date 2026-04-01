package fuzzers

import (
	"math/rand/v2"
	"strings"
)

const randomStrLen int = 32
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const RandomStringFuzzerName string = "randstr"

type RandomStringFuzzer struct{}

func (r RandomStringFuzzer) GenerateData() string {
	var output strings.Builder
	for range randomStrLen {
		output.WriteByte(charset[rand.IntN(len(charset))])
	}

	return output.String()
}
