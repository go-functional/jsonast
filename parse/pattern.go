package parse

// Pattern is the interface to match a token or tokens
type Pattern interface {
	IsValid(tkn Token) bool
	IsOpen() bool
}

type singleTokenPattern struct {
	expected Token
	found    bool
}

func (s singleTokenPattern) IsValid(tkn Token) bool {
	return tkn == s.expected
}

func (s singleTokenPattern) IsOpen() bool {
	return !s.found
}

type zeroOrMoreIdenticalTokensPattern struct {
	expected Token
}

func (z zeroOrMoreIdenticalTokensPattern) IsValid(tkn Token) bool {
	return tkn == z.expected
}

func (z zeroOrMoreIdenticalTokensPattern) IsOpen() bool {
	return true
}

