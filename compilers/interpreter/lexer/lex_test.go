package main

import "testing"

func TestTokenizer(t *testing.T) {
	// t.Fatal("not implemented")
	tests := []struct {
		input    string
		expected []token
	}{{
		input: ` let x = 2 + 45 * (8 - 4);`,
		expected: []token{
			token{Let, ""},
		},
	}}
	for _, tt := range tests {
		tokenizer := newTokenizer(tt.input)

		for _, expected := range tt.expected {
			actual := tokenizer.Next()
			if actual.Type != expected.Type {
				t.Fatal("fail")
			}
		}

	}
}
