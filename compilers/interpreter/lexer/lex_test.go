package lexer

import (
	"testing"

	"github.com/sthaha/interpreter/token"
)

func TestTokenizer(t *testing.T) {
	tests := []struct {
		input    string
		expected []token.Token
	}{{
		input: `=+(){}`,
		expected: []token.Token{
			tkn(token.Eq, '='),
			tkn(token.Plus, '+'),
			tkn(token.LParen, '('),
			tkn(token.RParen, ')'),
			tkn(token.LBrace, '{'),
			tkn(token.RBrace, '}'),
		},
	}}
	for _, tt := range tests {
		tokenizer := newTokenizer(tt.input)

		for _, expected := range tt.expected {
			actual := tokenizer.Next()
			if actual.Type != expected.Type {
				t.Fatalf("type error expected: %s | got: %s", expected.Type, actual.Type)
			}
			if actual.Value != expected.Value {
				t.Fatalf("value error expected: %q | got: %q", expected.Value, actual.Value)
			}
		}
	}
}
