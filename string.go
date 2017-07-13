package jsonast

import (
	"errors"
	"fmt"
)

var errNotAString = errors.New("not a string")

// String is a string JSON value
type String interface {
	Value
	fmt.Stringer
}

type stringImpl struct {
	Value
	str string
}

func (s stringImpl) String() string {
	return s.str
}

func isString(i interface{}) bool {
	_, ok := i.(string)
	return ok
}

func newString(v *value) (String, error) {
	if !isString(v.cur) {
		return nil, errNotAString
	}
	return &stringImpl{Value: v, str: v.cur.(string)}, nil
}
