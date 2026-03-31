package cli

import "fmt"

var ErrInvalidOutputType = fmt.Errorf("invalid output type. Valid output types are: 'csv', 'sql', 'json'")

type OutputType string

const (
	OutputTypeCSV  OutputType = "csv"
	OutputTypeSQL  OutputType = "sql"
	OutputTypeJSON OutputType = "json"
)

const DefaultOutputType OutputType = OutputTypeCSV

func (o OutputType) String() string {
	return string(o)
}

func (o *OutputType) Set(value string) error {
	if value != OutputTypeCSV.String() && value != OutputTypeJSON.String() && value != OutputTypeSQL.String() {
		return ErrInvalidOutputType
	}

	*o = OutputType(value)

	return nil
}
