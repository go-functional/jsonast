package jsonast

import (
	"errors"
	"fmt"
)

var errNotAnArray = errors.New("not an array")

type errIndexOutOfBounds struct {
	i int
}

func (e errIndexOutOfBounds) Error() string {
	return fmt.Sprintf("index %d is out of bounds", e.i)
}

// Array is an array JSON value
type Array interface {
	Value
	// Elts returns all of the elements in the array
	Elts() []Value
	// Foreach executes fn on each Value in Elts()
	Foreach(fn func(Value) error) error
	// At returns the Value at the given index i. Returns nil and
	// an appropriate error if len(Elts()) >= i
	At(i int) (Value, error)
}

type arrayImpl struct {
	Value
	arr []Value
}

func newArray(elts []Value) Array {
	return arrayImpl{
		Value: valueImpl{isArray: true},
		arr:   elts,
	}
}

func (v arrayImpl) Elts() []Value {
	return v.arr
}

func (v arrayImpl) Foreach(fn func(Value) error) error {
	for _, val := range v.arr {
		if err := fn(val); err != nil {
			return err
		}
	}
	return nil
}

func (v arrayImpl) At(i int) (Value, error) {
	if i >= len(v.arr) {
		return nil, errIndexOutOfBounds{i: i}
	}
	return v.arr[i], nil
}
