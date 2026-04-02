package fuzzers

import (
	"math/rand/v2"
	"strings"

	"github.com/grqphical/dumfuzz/pkg/libdumfuzz/dictionary"
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

const EmailFuzzerName string = "email"

var emailTLDS []string = []string{".com", ".ca", ".net", ".me", ".co", ".xyz", ".io"}

type EmailFuzzer struct{}

func (e EmailFuzzer) GenerateData() string {
	var output strings.Builder

	name := dictionary.GetRandomWordOfLength(14)
	output.WriteString(name)

	output.WriteString("@")

	domain := dictionary.GetRandomWordOfLength(6)
	output.WriteString(domain)

	output.WriteString(emailTLDS[rand.IntN(len(emailTLDS))])

	return output.String()
}
