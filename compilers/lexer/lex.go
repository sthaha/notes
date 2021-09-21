package main

type TOKEN string

var (
	EOF = TOKEN("EOF")
	LET = TOKEN("let")
)

type token struct {
	Type  TOKEN
	Value string
}
type tokenizer struct {
	input string
}

func newTokenizer(input string) *tokenizer {
	return &tokenizer{
		input: input,
	}
}

func (t *tokenizer) Next() *token {
	return &token{LET, ""}
}
