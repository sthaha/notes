package main

import "testing"

func TestTokenizer(t *testing.T) {
	// t.Fatal("not implemented")
	input := `
	let x = 2 + 45 * (8 - 4);
	`
	tokenizer := newTokenizer(input)
	token := tokenizer.Next()
	if token.Type != LET {
		t.Fatal("fail")
	}
}
