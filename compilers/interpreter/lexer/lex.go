package lexer

import (
	"github.com/sthaha/interpreter/token"
)

type lexer struct {
	input string
	next  int
	size  int
}

func New(input string) *lexer {
	return &lexer{
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

func (l *lexer) nextChar() (byte, bool) {
	ch, done := l.read()
	for isWhitespace(ch) {
		ch, done = l.read()
	}
	return ch, done
}

func (l *lexer) Next() token.Token {

	ch, done := l.nextChar()
	if done {
		return eof
	}

	switch ch {
	case ';':
		return semicolon
	case ',':
		return coma
	case '=':
		if ch, done := l.peek(); !done && ch == '=' {
			l.read()
			return eq
		}
		return assign
	case '!':
		if ch, done := l.peek(); !done && ch == '=' {
			l.read()
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
		if ch, done := l.peek(); !done && ch == '=' {
			l.read()
			return ge
		}
		return gt
	case '<':
		if ch, done := l.peek(); !done && ch == '=' {
			l.read()
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
			ret.Value = l.readWord()
			ret.Type = lookupType(ret.Value)
			// log.Printf("looking up: %q -> %v \n", ret.Value, rel.Type)
			// log.Printf("returning : %v  \n", ret)
			return ret
		}
		if isDigit(ch) {
			num := l.readNumber()
			return integer(num)
		}
		return sym(token.Illegal, ch)
	}
}

func (l *lexer) readWord() string {
	prev := l.next - 1
	for {
		ch, done := l.read()
		if done || !isLetter(ch) {
			l.next--
			break
		}
	}
	return l.input[prev:l.next]
}

func (l *lexer) readNumber() string {
	prev := l.next - 1
	for {
		ch, done := l.read()
		if done || !isDigit(ch) {
			l.next--
			break
		}
	}
	return l.input[prev:l.next]
}

func (l *lexer) peek() (byte, bool) {
	if l.next >= l.size {
		return 0, true
	}

	return l.input[l.next], false

}

func (l *lexer) read() (byte, bool) {
	if l.next >= l.size {
		return 0, true
	}

	ch := l.input[l.next]
	l.next++
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
