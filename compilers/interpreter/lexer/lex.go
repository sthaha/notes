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

func tkn(t token.Type, v byte) token.Token {
	return token.Token{Type: t, Value: string(v)}
}

func (t *tokenizer) Next() token.Token {

	ch, eof := t.read()
	if eof {
		return tkn(token.Eof, ch)
	}

	switch ch {
	case '=':
		return tkn(token.Eq, ch)
	case '+':
		return tkn(token.Plus, ch)
	case '(':
		return tkn(token.LParen, ch)
	case ')':
		return tkn(token.RParen, ch)
	case '{':
		return tkn(token.LBrace, ch)
	case '}':
		return tkn(token.RBrace, ch)
	default:
		return tkn(token.Illegal, ch)
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
