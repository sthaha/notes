package main

type TokenType string

const (
	Eof     TokenType = "EOF"
	Illegal           = "ILLEGAL"

	Semicolon  = ";"
	Whitespace = " "

	Let        = "let"
	Identifier = "Identifier"

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
