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

func (t *tokenizer) Next() token.Token {

	ch, done := t.read()
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
		return tkn(token.Illegal)
	}
}

func (t *tokenizer) read() (byte, bool) {
	if t.next >= t.size {
		return 0, true
	}

	ch := t.input[t.next]
	t.next++
	return ch, false

}
