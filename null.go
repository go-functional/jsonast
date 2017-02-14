package jsonast

import (
	"errors"
)

var errNotANull = errors.New("not null")

// Null is a null JSON value
type Null Value

func newNull(v Value) (Null, error) {
	if !v.IsNull() {
		return nil, errNotANull
	}
	return Null(v), nil
}
