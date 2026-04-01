package formats

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

var ErrInvalidOutputFormat = fmt.Errorf("invalid output type. Valid output types are: 'csv', 'sql', 'json'")

type OutputFormat string

const (
	OutputFormatCSV  OutputFormat = "csv"
	OutputFormatSQL  OutputFormat = "sql"
	OutputFormatJSON OutputFormat = "json"
)

const DefaultOutputFormat OutputFormat = OutputFormatCSV

func (o OutputFormat) String() string {
	return string(o)
}

func (o *OutputFormat) Set(value string) error {
	if value != OutputFormatCSV.String() && value != OutputFormatJSON.String() && value != OutputFormatSQL.String() {
		return ErrInvalidOutputFormat
	}

	*o = OutputFormat(value)

	return nil
}

func FormatCSV(data map[string][]string, out io.Writer) error {
	writer := csv.NewWriter(out)
	defer writer.Flush()

	keys := make([]string, 0, len(data))
	maxRows := 0
	for k, v := range data {
		keys = append(keys, k)
		if len(v) > maxRows {
			maxRows = len(v)
		}
	}
	sort.Strings(keys)

	if err := writer.Write(keys); err != nil {
		return err
	}

	for i := 0; i < maxRows; i++ {
		row := make([]string, len(keys))
		for j, key := range keys {
			// If this column has a value at this index, use it
			// Otherwise, leave it as an empty string (standard CSV behavior)
			if i < len(data[key]) {
				row[j] = data[key][i]
			} else {
				row[j] = ""
			}
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
