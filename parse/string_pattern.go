package parse

type StringPattern struct {
	openQuote  bool
	closeQuote bool
	str        string
}

func (s *StringPattern) IsValid(tkn Token) bool {
	if !s.openQuote && tkn == QuoteToken() {
		// no open quote yet and we found one means string start
		s.openQuote = true
		s.str += tkn.Char
		return true
	}
	if !s.openQuote && tkn != QuoteToken() {
		// no open quote yet and we didn't see one means invalid string
		return false
	}
	if s.openQuote && tkn == QuoteToken() {
		// open quote found and we see another means closed quote
		s.closeQuote = true
		s.str += tkn.Char
		return true
	}
	s.str += tkn.Char
	return true
}

func (s *StringPattern) IsOpen() bool {
	// pattern is still open if we haven't found the close quote yet
	return !s.closeQuote
}
