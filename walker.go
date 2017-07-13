package jsonast

import (
	"errors"
)

// Walker is a utility for walking the JSON tree with a few function calls,
// and transforming the JSON that was walked into a type
type Walker struct {
	path   string
	values []Value
	err    error
}

// Elt picks out the ith element of the current array
func (w Walker) Elt(i int) Walker {
	if w.err != nil {
		return w
	}
	if !w.curValue().IsArray() {
		w.err = errors.New("element at %s isn't an array", w.path)
		return w
	}
	// TODO
	return w
}

// Field picks out the key s in the current dictionary
func (w Walker) Field(s string) Walker {
	if w.err != nil {
		return w
	}
	// TODO
	return w
}
func (w Walker) And() Walker {
	if w.err != nil {
		return w
	}
	// TODO
	return w
}

func (w Walker) Validate(fn func(Value) bool) Walker {
	if w.err != nil {
		return w
	}
	fn(values)
}

func (w Walker) curValue() Value {
	// TODO: handle the case when there are no values
	return w.values[len(w.values)-1]
}
