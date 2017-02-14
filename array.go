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
	Elts() []Value
	At(i int) (Value, error)
}

type arrayImpl struct {
	Value
	arr []Value
}

func isArray(i interface{}) bool {
	_, ok := i.([]interface{})
	return ok
}

func newArray(v *value) (Array, error) {
	if !isArray(v.cur) {
		return nil, errNotAnArray
	}
	ifaceArr := v.cur.([]interface{})
	convertedArr := make([]Value, len(ifaceArr))
	for i, iface := range ifaceArr {
		val, err := valueFromIface(iface)
		if err != nil {
			return nil, err
		}
		convertedArr[i] = val
	}
	return &arrayImpl{Value: v, arr: convertedArr}, nil
}

func (v *arrayImpl) Elts() []Value {
	return v.arr
}

func (v *arrayImpl) At(i int) (Value, error) {
	if i >= len(v.arr) {
		return nil, errIndexOutOfBounds{i: i}
	}
	return v.arr[i], nil
}
