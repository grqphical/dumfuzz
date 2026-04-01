package main

import (
	"flag"
	"fmt"

	"github.com/grqphical/dumfuzz/internal/cli"
	"github.com/grqphical/dumfuzz/pkg/libdumfuzz"
)

func main() {
	var outputType cli.OutputType = cli.DefaultOutputType
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

	fmt.Printf("output type: %s, params: %+v, count: %d\n", outputType, params, *count)
}
