package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/grqphical/dumfuzz/internal/cli"
	"github.com/grqphical/dumfuzz/internal/formats"
	"github.com/grqphical/dumfuzz/pkg/libdumfuzz"
)

func main() {
	var outputType formats.OutputFormat = formats.DefaultOutputFormat
	flag.Var(&outputType, "output-type", "Output format of data (csv, json, sql)")

	var params cli.Params
	flag.Var(&params, "param", "Specifies the name and type of data for a field in the outputted dummy data. Format is `-param name=type`")

	listTypes := flag.Bool("list-types", false, "Lists all available fuzzers")
	count := flag.Int("n", 1, "How many data entries to output")

	flag.Parse()

	if *listTypes {
		for name := range libdumfuzz.Fuzzers {
			fmt.Println(name)
		}
		return
	}

	var generatedData map[string][]string = make(map[string][]string)

	for paramName, paramType := range params {
		data := libdumfuzz.RunFuzzer(libdumfuzz.GetBundledFuzzer(paramType), *count)
		generatedData[paramName] = data
	}

	switch outputType {
	case formats.OutputFormatCSV:
		formats.FormatCSV(generatedData, os.Stdout)
	}
}
