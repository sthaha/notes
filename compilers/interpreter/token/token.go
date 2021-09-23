package token

type Type string

const (
	Eof     Type = "EOF"
	Illegal      = "ILLEGAL"

	Semicolon  = ";"
	Whitespace = " "

	Let        = "let"
	Identifier = "Identifier"

	// operators
	Plus Type = "+"
	Sub       = "-"
	Mul       = "*"
	Div       = "/"
	Mod       = "%"

	Eq Type = "="

	LParen = "("
	RParen = ")"

	LBrace = "{"
	RBrace = "}"
)

type Token struct {
	Type  Type
	Value string
}
