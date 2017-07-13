package jsonast

import (
	"github.com/go-functional/jsonast/parse"
)

type state struct {
	inObj  bool
	inArr  bool
	inStr  bool
	curVal Value
	aggVal Value
}

func newState() state {
	return state{}
}

func (s state) addToken(tkn parse.Token) error {
	// char := tkn.char
	switch tkn {
	case parse.QuoteToken():
		if !s.inStr {
			// start a new string
		} else {
			// end a string
		}
	}
	return nil
}

func (s state) value() Value {
	// TODO
	return nil
}

// Parse parses jsonStr from JSON to a Value. Returns nil and an appropriate
// error if jsonStr was an invalid JSON string
func Parse(jsonStr string) (Value, error) {
	tokensCh := make(chan parse.Token)
	go parse.Tokenize(jsonStr, tokensCh)
	st := newState()
	for token := range tokensCh {
		if err := st.addToken(token); err != nil {
			return nil, err
		}
	}
	return st.value(), nil
}
