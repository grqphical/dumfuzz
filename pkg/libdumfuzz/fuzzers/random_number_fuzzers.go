package fuzzers

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

const RandomIntegerFuzzerName string = "randint"

type RandomIntegerFuzzer struct{}

func (r RandomIntegerFuzzer) GenerateData() string {
	return strconv.Itoa(rand.IntN(128 * 1024))
}

const RandomFloatFuzzerName string = "randfloat"

type RandomFloatFuzzer struct{}

func (r RandomFloatFuzzer) GenerateData() string {
	return strconv.FormatFloat(rand.Float64(), 'f', 4, 64)
}

const RandomBigIntegerFuzzerName string = "randbigint"

type RandomBigIntegerFuzzer struct{}

func (r RandomBigIntegerFuzzer) GenerateData() string {
	return strconv.Itoa(rand.Int())
}

const PhoneNumberFuzzerName string = "phone"

type PhoneNumberFuzzer struct{}

func (r PhoneNumberFuzzer) GenerateData() string {
	var output strings.Builder
	output.WriteByte('+')
	output.Write([]byte(strconv.Itoa(1 + rand.IntN(999)))) // generate country code

	// generate subscriber number (rest of phone number)
	for range 10 {
		output.WriteByte(uint8(48 + rand.IntN(57-48+1)))
	}

	return output.String()
}
