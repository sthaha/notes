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
	// keywords
	let      = tkn(token.Let)
	kwTrue   = tkn(token.True)
	kwFalse  = tkn(token.False)
	kwIf     = tkn(token.If)
	kwElse   = tkn(token.Else)
	kwReturn = tkn(token.Return)

	coma      = tkn(token.Coma)
	semicolon = tkn(token.Semicolon)
	assign    = tkn(token.Assign)
	// arithmetic operators
	plus     = tkn(token.Plus)
	minus    = tkn(token.Minus)
	asterisk = tkn(token.Asterisk)
	slash    = tkn(token.Slash)

	gt  = tkn(token.GT)
	eq  = tkn(token.Eq)
	neq = tkn(token.NE)
	ge  = tkn(token.GE)
	lt  = tkn(token.LT)
	le  = tkn(token.LE)

	bang = tkn(token.Bang)

	fn     = tkn(token.Function)
	lParen = tkn(token.LParen)
	rParen = tkn(token.RParen)
	lBrace = tkn(token.LBrace)
	rBrace = tkn(token.RBrace)
	eof    = sym(token.EOF, 0)
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
	case ';':
		return semicolon
	case ',':
		return coma
	case '=':
		if ch, done := t.peek(); !done && ch == '=' {
			t.read()
			return eq
		}
		return assign
	case '!':
		if ch, done := t.peek(); !done && ch == '=' {
			t.read()
			return neq
		}
		return bang

	case '+':
		return plus
	case '-':
		return minus
	case '*':
		return asterisk
	case '/':
		return slash

	case '>':
		if ch, done := t.peek(); !done && ch == '=' {
			t.read()
			return ge
		}
		return gt
	case '<':
		if ch, done := t.peek(); !done && ch == '=' {
			t.read()
			return le
		}
		return lt

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
		if isDigit(ch) {
			num := t.readNumber()
			return integer(num)
		}
		return sym(token.Illegal, ch)
	}
}

func (t *tokenizer) readWord() string {
	prev := t.next - 1
	for {
		ch, done := t.read()
		if done || !isLetter(ch) {
			t.next--
			break
		}
	}
	return t.input[prev:t.next]
}

func (t *tokenizer) readNumber() string {
	prev := t.next - 1
	for {
		ch, done := t.read()
		if done || !isDigit(ch) {
			t.next--
			break
		}
	}
	return t.input[prev:t.next]
}

func (t *tokenizer) peek() (byte, bool) {
	if t.next >= t.size {
		return 0, true
	}

	return t.input[t.next], false

}

func (t *tokenizer) read() (byte, bool) {
	if t.next >= t.size {
		return 0, true
	}

	ch := t.input[t.next]
	t.next++
	return ch, false

}

var (
	keywords = map[string]token.Type{
		"let":    token.Let,
		"fn":     token.Function,
		"true":   token.True,
		"false":  token.False,
		"if":     token.If,
		"else":   token.Else,
		"return": token.Return,
	}
)

func lookupType(x string) token.Type {
	if t, ok := keywords[x]; ok {
		return t
	}
	return token.Identifier
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
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
