package cli

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidKeyValuePair = errors.New("invalid key-value argument")
	ErrEmptyKey            = errors.New("key cannot be empty")
)

// custom wrapper type to parse CLI parameters for dummy data fuzzer
type Params map[string]string

func (p *Params) String() string {
	if p == nil || *p == nil {
		return "{}"
	}
	return fmt.Sprintf("%v", map[string]string(*p))
}

func (p *Params) Set(value string) error {
	if *p == nil {
		*p = make(Params)
	}

	pair := strings.SplitN(value, "=", 2)
	if len(pair) != 2 {
		return ErrInvalidKeyValuePair
	}

	key := strings.TrimSpace(pair[0])
	val := strings.TrimSpace(pair[1])

	if key == "" {
		return ErrEmptyKey
	}
	(*p)[key] = val
	return nil
}
