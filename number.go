package jsonast

import (
	"errors"
)

var errNotANumber = errors.New("not a number")

// Number is the numeric JSON value
type Number interface{}

type number struct {
	Value
	i float64
}

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
	default:
		return -1, errNotANumber

	}
}

func isNumber(i interface{}) bool {
	_, err := convertNumber(i)
	return err == nil
}

func newNumber(v *value) (Number, error) {
	flt, err := convertNumber(v.cur)
	if err != nil {
		return nil, err
	}
	return &number{Value: v, i: flt}, nil
}
