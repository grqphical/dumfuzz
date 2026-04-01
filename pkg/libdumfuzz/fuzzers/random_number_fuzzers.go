package fuzzers

import (
	"math/rand/v2"
	"strconv"
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
