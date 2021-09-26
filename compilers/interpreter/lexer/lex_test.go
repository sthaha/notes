package lexer

import (
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

func TestTokenizer_All(t *testing.T) {

	input := `
		let five = 5;
		let ten = 10;

		let add = fn(x, y) {
			x + y;
		};

		let result = add(five, ten);
		!-/*+5;
		5 > 3;
		3 < 5;
		!true == false
		-5
		10 == 10;
		 0 != 10;
	`
	expected := []token.Token{
		let, id("five"), assign, integer("5"), semicolon,
		let, id("ten"), assign, integer("10"), semicolon,

		let, id("add"), assign, fn, lParen, id("x"), coma, id("y"), rParen, lBrace,
		id("x"), plus, id("y"), semicolon,
		rBrace, semicolon,

		let, id("result"), assign,
		id("add"), lParen, id("five"), coma, id("ten"), rParen, semicolon,
		bang, minus, slash, asterisk, plus, integer("5"), semicolon,
		integer("5"), gt, integer("3"), semicolon,
		integer("3"), lt, integer("5"), semicolon,
		eof,
	}

	assertTokens(t, input, expected)

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
		id("x"), plus, id("y"), semicolon,
		rBrace, semicolon,

		let, id("result"), assign,
		id("add"), lParen, id("five"), coma, id("ten"), rParen, semicolon,
	}

	assertTokens(t, input, expected)

}
func assertToken(t *testing.T, expected, actual token.Token) {
	if actual.Type != expected.Type {
		t.Fatalf("type error expected: %s | got: %s", expected, actual)
	}
	if actual.Value != expected.Value {
		t.Fatalf("value error expected: %q | got: %q", expected.Value, actual.Value)
	}
}

func assertTokens(t *testing.T, input string, expected []token.Token) {
	tokenizer := newTokenizer(input)

	for _, expected := range expected {
		actual := tokenizer.Next()
		assertToken(t, expected, actual)
	}

}
