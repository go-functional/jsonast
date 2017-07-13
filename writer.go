package jsonast

import (
	"fmt"
	"io"
)

// Write writes val as JSON to w
func Write(val Value, w io.Writer) error {
	switch x := val.(type) {
	case String:
		_, err := fmt.Fprintf(w, `"%s"`, x.String())
		return err
	case Number:
		_, err := fmt.Fprint(w, x.Float64())
		return err
	default:
		return fmt.Errorf("unknown value found: %#v", x)
	}
	return nil
}
