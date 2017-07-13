package jsonast

import (
	"bytes"
	"testing"

	"github.com/arschles/assert"
)

func TestWriteString(t *testing.T) {
	const (
		rawStr  = "abc"
		rawJSON = `"` + rawStr + `"`
	)
	val := &value{b: []byte(rawJSON), cur: rawStr}
	str, err := newString(val)
	assert.NoErr(t, err)
	buf := bytes.NewBuffer(nil)
	assert.NoErr(t, Write(str, buf))
	retJSON := string(buf.Bytes())
	assert.Equal(t, retJSON, rawJSON, "returned JSON")
}
