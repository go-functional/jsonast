package jsonast

import (
	"encoding/json"
)

// Value is a generic node in a JSON AST. This implements fmt.Stringer
type Value interface {
	// Walk creates a walker, which is capable of walking the JSON
	// tree with a few function calls
	Walk() Walker
	// IsString returns true if the data at this node is a string
	IsString() bool
	// IsNumber returns true if the data at this node is numeric
	IsNumber() bool
	// IsBool returns true if the data at this node is a bool
	IsBool() bool
	// IsObject returns true if the data at this node is an object
	IsObject() bool
	// IsArray returns true if the data at this node is an array
	IsArray() bool
	// IsNull returns true if the data at this node is null
	IsNull() bool
}

type value struct {
	b   []byte
	cur interface{}
}

// ValueFromBytes decodes b into a Value, or returns an error if b is invalid
// json
func ValueFromBytes(b []byte) (Value, error) {
	var iface interface{}
	if err := json.Unmarshal(b, &iface); err != nil {
		return nil, err
	}
	return &value{cur: iface, b: b}, nil
}

func valueFromIface(i interface{}) (Value, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	return &value{b: b, cur: i}, nil
}

func (v *value) Bytes() []byte {
	return v.b
}

// func (v *value) String() string {
// 	return string(v.b)
// }

func (v *value) IsString() bool {
	return isString(v.cur)
}

// func (v *value) StringVal() (String, error) {
// 	return newString(v)
// }

func (v *value) IsNumber() bool {
	return isNumber(v.cur)
}

// func (v *value) NumberVal() (Number, error) {
// 	return newNumber(v)
// }

func (v *value) IsBool() bool {
	return isBool(v)
}

// func (v *value) BoolVal() (Bool, error) {
// 	if !v.IsBool() {
// 		return nil, errNotABool
// 	}
// 	return newBool(v), nil
// }

func (v *value) IsObject() bool {
	_, ok := v.cur.(map[interface{}]interface{})
	return ok
}

// func (v *value) ObjectVal() (Object, error) {
// 	if !v.IsObject() {
// 		return nil, errNotAnObject
// 	}
// 	return newObject(v)
// }

func (v *value) IsArray() bool {
	return isArray(v.cur)
}

// func (v *value) ArrayVal() (Array, error) {
// 	if !v.IsArray() {
// 		return nil, errNotAnArray
// 	}
// 	return newArray(v)
// }

func (v *value) IsNull() bool {
	return v.cur == nil
}

// func (v *value) NullVal() (Null, error) {
// 	return newNull(v)
// }
