package main

import (
	"flag"
	"fmt"

	"github.com/grqphical/dumfuzz/internal/cli"
)

func main() {
	var outputType cli.OutputType = cli.DefaultOutputType
	flag.Var(&outputType, "output-type", "Output format of data (csv, json, sql)")

	flag.Parse()

	fmt.Printf("output type: %s\n", outputType)
}
