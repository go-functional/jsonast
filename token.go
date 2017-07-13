package jsonast

type token struct {
	char string
}

func (t token) isDigit() bool {
	r := t.char
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

var commaToken = token{char: `,`}
var quoteToken = token{char: `"`}
var arrayStartToken = token{char: `[`}
var arrayEndToken = token{char: `]`}
var objectStartToken = token{char: `{`}
var objectEndToken = token{char: `}`}

func tokenize(str string, ch chan<- token) {
	for _, char := range str {
		var tkn token
		if char == ',' {
			tkn = commaToken
		} else if char == '"' {
			tkn = quoteToken
		} else if char == '[' {
			tkn = arrayStartToken
		} else if char == ']' {
			tkn = arrayEndToken
		} else if char == '{' {
			tkn = objectStartToken
		} else if char == '}' {
			tkn = objectEndToken
		} else {
			tkn = token{char: char}
		}
		ch <- tkn
	}
	close(ch)
}
