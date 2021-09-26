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

func TestTokenizer_Keywords(t *testing.T) {

	input := `
    if (5 < 10 ) {
      return true;
    } else {
      return false;
    }
	`
	expected := []token.Token{
		kwIf, lParen, integer("5"), lt, integer("10"), rParen, lBrace,
		kwReturn, kwTrue, semicolon,
		rBrace, kwElse, lBrace,
		kwReturn, kwFalse, semicolon,
		rBrace,
		eof,
	}

	assertTokens(t, input, expected)
}

func TestTokenizer_Arithmetic(t *testing.T) {

	input := `
		!-/*+5;
		5 > 3;
		3 < 5;
		3 < 10 > 5;
	`
	expected := []token.Token{
		bang, minus, slash, asterisk, plus, integer("5"), semicolon,
		integer("5"), gt, integer("3"), semicolon,
		integer("3"), lt, integer("5"), semicolon,
		integer("3"), lt, integer("10"), gt, integer("5"), semicolon,
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
