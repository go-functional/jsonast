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

func newBool(b bool) Bool {
	return boolImpl{
		Value: valueImpl{isBool: true},
		val:   b,
	}
}

func (b boolImpl) True() bool {
	return b.val
}
