package lexer

import (
	"log"
	"testing"

	"github.com/sthaha/interpreter/token"
)

func TestTokenizer_Simple(t *testing.T) {
	input := `=+(){}`
	expected := []token.Token{assign, plus, lParen, rParen, lBrace, rBrace, eof}
	assertTokens(t, input, expected)
}
func TestTokenizer_id(t *testing.T) {
	// let, id("five"), assign, integer("5"), semicolon
	assertToken(t, token.Token{Type: token.Identifier, Value: "five"}, id("five"))
}

func TestTokenizer_Helloworld(t *testing.T) {

	input := `
    let five = 5;
    let ten = 10;

    let add = fn(x, y) {
      x + y;
    };

    let result = add(five, ten);
	`
	expected := []token.Token{
		let, id("five"), assign, integer("5"), semicolon,
		let, id("ten"), assign, integer("10"), semicolon,

		let, id("add"), assign, fn, lParen, id("x"), coma, id("y"), rParen, lBrace,
		id("x"), plus, id("y"),
		rBrace, semicolon,
	}

	assertTokens(t, input, expected)

}
func assertToken(t *testing.T, expected, actual token.Token) {
	if actual.Type != expected.Type {
		t.Fatalf("type error expected: %s | got: %s", expected.Type, actual.Type)
	}
	if actual.Value != expected.Value {
		t.Fatalf("value error expected: %q | got: %q", expected.Value, actual.Value)
	}
}

func assertTokens(t *testing.T, input string, expected []token.Token) {
	tokenizer := newTokenizer(input)

	for _, expected := range expected {
		actual := tokenizer.Next()
		log.Printf("   ... %v   |   %v", actual, expected)
		assertToken(t, expected, actual)
	}

}
