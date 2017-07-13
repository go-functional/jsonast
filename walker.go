package jsonast

import (
	"fmt"
)

// Walker is a utility for walking the JSON tree with a few function calls,
// and transforming the JSON that was walked into a type
type Walker struct {
	initialValue Value
	latestPath   string
	latestValue  Value
	Nulls        []Null
	Strings      []String
	Numbers      []Number
	Objects      []Object
	Arrays       []Array
	err          error
}

// NewWalker creates a new Walker from the given initVal
func NewWalker(initVal Value) Walker {
	return Walker{initialValue: initVal}
}

// Elt picks out the ith element of the current array
func (w Walker) Elt(i int) Walker {
	if w.err != nil {
		return w
	}
	if !w.latestValue.IsArray() {
		w.err = fmt.Errorf("element at %s isn't an array", w.latestPath)
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
	return Walker{}
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
	if !fn(w.latestValue) {
		w.err = fmt.Errorf("validation on path %s failed", w.latestPath)
		return w
	}
	return w
}
