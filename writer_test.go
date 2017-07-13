package jsonast

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/arschles/assert"
)

type writerTestCase struct {
	rawVal     interface{}
	expectJSON string
	newVal     func(v *value) (Value, error)
}

func (w writerTestCase) run() error {
	rawVal := &value{b: []byte(w.expectJSON), cur: w.rawVal}
	val, err := w.newVal(rawVal)
	if err != nil {
		return fmt.Errorf("error calling newVal (%s)", err)
	}
	buf := bytes.NewBuffer(nil)
	if err := Write(val, buf); err != nil {
		return fmt.Errorf("error writing (%s)", err)
	}
	retJSON := string(buf.Bytes())
	if retJSON != w.expectJSON {
		return fmt.Errorf("expected json '%s', got '%s'", w.expectJSON, retJSON)
	}
	return nil

}

func TestWriteString(t *testing.T) {
	tc := writerTestCase{
		rawVal:     "abc",
		expectJSON: `"abc"`,
		newVal: func(v *value) (Value, error) {
			str, err := newString(v)
			if err != nil {
				return nil, err
			}
			return str, nil
		},
	}
	assert.NoErr(t, tc.run())
}

func TestWriteNumber(t *testing.T) {
	tc := writerTestCase{
		rawVal:     123.401,
		expectJSON: `123.401`,
		newVal: func(v *value) (Value, error) {
			num, err := newNumber(v)
			if err != nil {
				return nil, err
			}
			return num, nil
		},
	}
	assert.NoErr(t, tc.run())
}
