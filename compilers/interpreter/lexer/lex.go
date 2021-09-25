package lexer

import (
	"github.com/sthaha/interpreter/token"
)

type tokenizer struct {
	input string
	next  int
	size  int
}

func newTokenizer(input string) *tokenizer {
	return &tokenizer{
		input: input,
		next:  0,
		size:  len(input),
	}
}

func tkn(t token.Type) token.Token {
	return token.Token{Type: t, Value: string(t)}
}

func id(name string) token.Token {
	return token.Token{Type: token.Identifier, Value: name}
}

func integer(i string) token.Token {
	return token.Token{Type: token.Identifier, Value: i}
}

func sym(t token.Type, b byte) token.Token {
	return token.Token{Type: t, Value: string(b)}
}

var (
	let       = tkn(token.Let)
	assign    = tkn(token.Assign)
	coma      = tkn(token.Coma)
	semicolon = tkn(token.Semicolon)
	plus      = tkn(token.Plus)
	fn        = tkn(token.Function)
	lParen    = tkn(token.LParen)
	rParen    = tkn(token.RParen)
	lBrace    = tkn(token.LBrace)
	rBrace    = tkn(token.RBrace)
	eof       = sym(token.EOF, 0)
)

func (t *tokenizer) nextChar() (byte, bool) {
	ch, done := t.read()
	for isWhitespace(ch) {
		ch, done = t.read()
	}
	return ch, done
}

func (t *tokenizer) Next() token.Token {

	ch, done := t.nextChar()
	if done {
		return eof
	}

	switch ch {
	case '=':
		return assign
	case '+':
		return plus
	case '(':
		return lParen
	case ')':
		return rParen
	case '{':
		return lBrace
	case '}':
		return rBrace
	default:

		if isLetter(ch) {
			ret := token.Token{}
			ret.Value = t.readWord()
			ret.Type = lookupType(ret.Value)
			// log.Printf("looking up: %q -> %v \n", ret.Value, ret.Type)
			// log.Printf("returning : %v  \n", ret)
			return ret
		}
		// if isDigit(ch) {

		// }
		return sym(token.Illegal, ch)
	}
}

func (t *tokenizer) readWord() string {
	prev := t.next - 1
	for {
		ch, done := t.read()
		if done || !isLetter(ch) {
			break
		}
	}
	return t.input[prev : t.next-1]
}

var (
	keywords = map[string]token.Type{
		"let": token.Let,
		"fn":  token.Function,
	}
)

func lookupType(x string) token.Type {
	if t, ok := keywords[x]; ok {
		return t
	}
	return token.Identifier
}

func isLetter(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z')
}

func isWhitespace(b byte) bool {
	switch b {
	case ' ', '\t', '\n', '\r':
		return true
	}
	return false
}

func (t *tokenizer) read() (byte, bool) {
	if t.next >= t.size {
		return 0, true
	}

	ch := t.input[t.next]
	t.next++
	return ch, false

}
