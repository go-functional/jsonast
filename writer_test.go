package jsonast

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/arschles/assert"
)

type writerTestCase struct {
	val        Value
	expectJSON string
}

func (w writerTestCase) run() error {
	buf := bytes.NewBuffer(nil)
	if err := Write(w.val, buf); err != nil {
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
		val:        newString("abc"),
		expectJSON: `"abc"`,
	}
	assert.NoErr(t, tc.run())
}

func TestWriteNumber(t *testing.T) {
	tc := writerTestCase{
		val:        newNumber(123.401),
		expectJSON: `123.401`,
	}
	assert.NoErr(t, tc.run())
}
