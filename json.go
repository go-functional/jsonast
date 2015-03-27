package jsonast

import (
	"encoding/json"
	"reflect"
)

type JSON struct {
	root interface{}
}

func Decode(b []byte) (*JSON, error) {
	var i interface{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}
	return &JSON{root: i}, nil
}

func (j *JSON) IsInt() bool {
	_, ok := j.root.(int)
	return ok
}

func (j *JSON) Int() int {
	if !j.IsInt() {
		return 0
	}
	return j.root.(int)
}

func (j *JSON) IsBool() bool {
	_, ok := j.root.(bool)
	return ok
}

func (j *JSON) Bool() bool {
	if !j.IsBool() {
		return false
	}
	return j.root.(bool)
}

func (j *JSON) IsArr() bool {
	return reflect.ValueOf(j.root).Kind() == reflect.Array
}
