package jsonast

import (
	"log"
)

type state struct {
	curStr string

	inObj  bool
	inArr  bool
	inStr  bool
	curVal Value
	aggVal Value
}

func (s state) addToken(tkn token) state {
	char := tkn.char
	switch tkn {
	case commaToken:
		if s.inStr {
			s.curStr += char
		}
		// TODO: if in array, not in string
	case quoteToken:
	case commaToken:

	}
	return s
}

// Parse parses jsonStr from JSON to a Value. Returns nil and an appropriate
// error if jsonStr was an invalid JSON string
func Parse(jsonStr string) (Value, error) {
	tokensCh := make(chan token)
	go tokenize(jsonStr, tokensCh)
	var st state
	for token := range tokensCh {
		st = st.addChar(token)
		log.Printf("found token %s", token)
	}
	return nil, nil
}
