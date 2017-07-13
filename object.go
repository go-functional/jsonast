package jsonast

import (
	"errors"
	"fmt"
)

var (
	errNotAnObject = errors.New("not an object")
)

// Object is the object JSON value
type Object interface {
	Value
	Fields() []ObjectPair
	ValueAt(key string) (Value, error)
}

// ObjectPair represents a single key-value pair in an object
type ObjectPair struct {
	Key   string
	Value Value
}

type objectImpl struct {
	Value
	pairs []ObjectPair
}

func newObject(pairs []ObjectPair) Object {
	return objectImpl{
		Value: valueImpl{isObject: true},
		pairs: pairs,
	}
}

func (o objectImpl) Fields() []ObjectPair {
	return o.pairs
}

func (o objectImpl) ValueAt(key string) (Value, error) {
	for _, pair := range o.pairs {
		if pair.Key == key {
			return pair.Value, nil
		}
	}
	return nil, fmt.Errorf("no such key %s", key)
}
