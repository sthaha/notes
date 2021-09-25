package token

type Type string

const (
	EOF     Type = "EOF"
	Illegal      = "ILLEGAL"

	Semicolon Type = ";"
	Coma           = ","

	Let        Type = "let"
	Identifier      = "Identifier"

	// operators
	Plus Type = "+"
	Sub       = "-"
	Mul       = "*"
	Div       = "/"
	Mod       = "%"

	Assign Type = "="

	LParen Type = "("
	RParen      = ")"
	LBrace      = "{"
	RBrace      = "}"

	// datatypes
	Integer Type = "integer"
	String       = "string"

	Function Type = "fn"
)

type Token struct {
	Type  Type
	Value string
}
