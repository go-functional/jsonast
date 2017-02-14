package jsonast

import (
	"errors"
)

var errNotAString = errors.New("not a string")

// String is a string JSON value
type String interface {
	Value
}

type stringImpl struct {
	Value
}

func isString(i interface{}) bool {
	_, ok := i.(string)
	return ok
}

func newString(v *value) (String, error) {
	if !isString(v.cur) {
		return nil, errNotAString
	}
	return &stringImpl{Value: v}, nil
}
