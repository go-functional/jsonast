package jsonast

import (
	"fmt"
)

// PathElt is a container for a single element of a json path
type PathElt interface {
	fmt.Stringer
}

type stringPathElt struct {
	str string
}

func (s stringPathElt) String() string {
	return fmt.Sprintf("StringPathElt(%s)", s.str)
}

type intPathElt struct {
	i int
}

func (i intPathElt) String() string {
	return fmt.Sprintf("IntPathElt(%d)", i)
}

// NewStringPathElt creates a new PathElt from s
func NewStringPathElt(s string) PathElt {
	return stringPathElt{str: s}
}

// NewIntPathElt creates a new PathElt from i
func NewIntPathElt(i int) PathElt {
	return intPathElt{i: i}
}
