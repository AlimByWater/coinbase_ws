package proto

import (
	"errors"
	"strings"
)

type (
	Symbol string
)

func (s *Symbol) Validate() error {
	if len(string(*s)) <= 1 && !strings.ContainsAny(string(*s), "-") {
		return errors.New("Invalid symbol pair: " + string(*s))
	}

	return nil
}
