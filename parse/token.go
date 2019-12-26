package parse

// Token represents a single token in the parser
type Token struct {
	Char string
}

func (t Token) IsDigit() bool {
	r := t.Char
	return r == "0" ||
		r == "1" ||
		r == "2" ||
		r == "3" ||
		r == "4" ||
		r == "5" ||
		r == "6" ||
		r == "7" ||
		r == "8" ||
		r == "9"
}

func CommaToken() Token {
	return Token{Char: `,`}
}
func QuoteToken() Token {
	return Token{Char: `"`}
}

func ArrayStartToken() Token {
	return Token{Char: `[`}
}

func ArrayEndToken() Token {
	return Token{Char: `]`}
}

func ObjectStartToken() Token {
	return Token{Char: `{`}
}
func ObjectEndToken() Token {
	return Token{Char: `}`}
}

func Tokenize(str string, ch chan<- Token) {
	for _, char := range str {
		var tkn Token
		if char == ',' {
			tkn = CommaToken()
		} else if char == '"' {
			tkn = QuoteToken()
		} else if char == '[' {
			tkn = ArrayStartToken()
		} else if char == ']' {
			tkn = ArrayEndToken()
		} else if char == '{' {
			tkn = ObjectStartToken()
		} else if char == '}' {
			tkn = ObjectEndToken()
		} else {
			tkn = Token{Char: string(char)}
		}
		ch <- tkn
	}
	close(ch)
}
