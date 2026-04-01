package libdumfuzz

import "github.com/grqphical/dumfuzz/pkg/libdumfuzz/fuzzers"

type DummyDataFuzzer interface {
	GenerateData() string
}

var Fuzzers map[string]DummyDataFuzzer = make(map[string]DummyDataFuzzer)

func init() {
	Fuzzers[fuzzers.RandomIntegerFuzzerName] = fuzzers.RandomIntegerFuzzer{}
	Fuzzers[fuzzers.RandomFloatFuzzerName] = fuzzers.RandomFloatFuzzer{}
}
