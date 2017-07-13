package jsonast

import (
	"errors"
)

// Walker is a utility for walking the JSON tree with a few function calls,
// and transforming the JSON that was walked into a type
type Walker struct {
	latestValue Value
	Nulls       []Null
	Strings     []String
	Numbers     []Number
	Objects     []Object
	Arrays      []Array
	err         error
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

// Path picks out the value pointed to by all the given path elements
// and puts it into the next value
func (w Walker) Path(pathElt PathElt, pathElts ...PathElt) Walker {
	if w.err != nil {
		return w
	}
	// TODO
	return nil
}

// Field picks out the key s in the current dictionary
func (w Walker) Field(s string) Walker {
	if w.err != nil {
		return w
	}
	return w.Path(NewStringPathElt(s))
}

// Validate validates the latest fetched value in w
func (w Walker) Validate(fn func(Value) bool) Walker {
	if w.err != nil {
		return w
	}
	fn(values)
}
