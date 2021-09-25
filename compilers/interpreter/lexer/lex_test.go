package lexer

import (
	"testing"

	"github.com/sthaha/interpreter/token"
)

func TestTokenizer_Simple(t *testing.T) {
	input := `=+(){}`
	expected := []token.Token{
		tkn(token.Eq, '='),
		tkn(token.Plus, '+'),
		tkn(token.LParen, '('),
		tkn(token.RParen, ')'),
		tkn(token.LBrace, '{'),
		tkn(token.RBrace, '}'),
	}
	assertTokens(t, input, expected)
}

func assertTokens(t *testing.T, input string, expected []token.Token) {
	tokenizer := newTokenizer(input)

	for _, expected := range expected {
		actual := tokenizer.Next()
		if actual.Type != expected.Type {
			t.Fatalf("type error expected: %s | got: %s", expected.Type, actual.Type)
		}
		if actual.Value != expected.Value {
			t.Fatalf("value error expected: %q | got: %q", expected.Value, actual.Value)
		}
	}

}

//
//
// “input := `let five = 5;
// let ten = 10;
//
// let add = fn(x, y) {
//   x + y;
// };
//
// let result = add(five, ten);
// `
// ”
//
