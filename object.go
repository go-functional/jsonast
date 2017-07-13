package jsonast

import (
	"errors"
)

var (
	errNotAnObject = errors.New("not an object")
)

// Object is the object JSON value
type Object interface {
	Value
	Fields() []ObjectPair
	ValueAt(name string) (Value, error)
}

// ObjectPair represents a single key-value pair in an object
type ObjectPair struct {
	Name  string
	Value Value
}

type object struct {
	Value
	obj []ObjectPair
}

func isObject(i interface{}) bool {
	_, ok := i.(map[interface{}]interface{})
	return ok
}
