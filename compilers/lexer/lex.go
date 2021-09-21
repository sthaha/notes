package main

type TokenType string

const (
	Eof     TokenType = "EOF"
	Illegal           = "ILLEGAL"

	Semicolon = ";"

	Let = "let"

	// operators
	Plus TokenType = "+"
	Sub            = "-"
	Mul            = "*"
	Div            = "/"
	Mod            = "%"

	Eq TokenType = "="

	LeftParan  = "("
	RightParan = ")"
)

type token struct {
	Type  TokenType
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
	return &token{Let, ""}
}
