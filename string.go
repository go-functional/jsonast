package jsonast

import (
	"errors"
	"fmt"
)

var errNotAString = errors.New("not a string")

// String is a string JSON value
type String interface {
	fmt.Stringer
	Value
}

type stringImpl struct {
	Value
	str string
}

func (s stringImpl) String() string {
	return s.str
}

func newString(str string) String {
	return stringImpl{
		Value: valueImpl{isString: true},
		str:   str,
	}
}
