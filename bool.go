package jsonast

import (
	"errors"
)

var errNotABool = errors.New("not a bool")

// Bool is a JSON bool type
type Bool interface {
	Value
	True() bool
}

type boolImpl struct {
	Value
	val bool
}

func isBool(i interface{}) bool {
	_, ok := i.(bool)
	return ok
}

func newBool(v *value) (Bool, error) {
	if !isBool(v.cur) {
		return nil, errNotABool
	}
	return &boolImpl{
		Value: v,
		val:   v.cur.(bool),
	}, nil
}

func (b *boolImpl) True() bool {
	return b.val
}
