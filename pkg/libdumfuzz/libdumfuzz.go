package libdumfuzz

import (
	"github.com/grqphical/dumfuzz/pkg/libdumfuzz/fuzzers"
)

type DummyDataFuzzer interface {
	GenerateData() string
}

var Fuzzers map[string]DummyDataFuzzer = make(map[string]DummyDataFuzzer)

// Looks up an included fuzzer by name, returns nil if no fuzzer with that name is found
func GetBundledFuzzer(name string) DummyDataFuzzer {
	fuzzer, exists := Fuzzers[name]
	if !exists {
		return nil
	}

	return fuzzer
}

// Runs the given fuzzer n times, returning an array of strings containing the generated data
func RunFuzzer(fuzzer DummyDataFuzzer, n int) []string {
	result := make([]string, n)

	for i := range n {
		result[i] = fuzzer.GenerateData()
	}

	return result
}

func init() {
	Fuzzers[fuzzers.RandomIntegerFuzzerName] = fuzzers.RandomIntegerFuzzer{}
	Fuzzers[fuzzers.RandomFloatFuzzerName] = fuzzers.RandomFloatFuzzer{}
}
