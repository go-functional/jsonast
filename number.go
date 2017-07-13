package jsonast

import (
	"errors"
)

var errNotANumber = errors.New("not a number")

// Number is the numeric JSON value
type Number interface {
	Value
	Float64() float64
}

type number struct {
	Value
	i float64
}

func newNumber(f float64) Number {
	return &number{
		Value: valueImpl{isNumber: true},
		i:     f,
	}
}

func (n *number) Float64() float64 {
	return n.i
}

// TODO: remove?
func convertNumber(i interface{}) (float64, error) {
	switch t := i.(type) {
	case int:
		return float64(t), nil
	case int8:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case uint8:
		return float64(t), nil
	case uint16:
		return float64(t), nil
	case uint32:
		return float64(t), nil
	case uint64:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return t, nil
	default:
		return -1, errNotANumber

	}
}
