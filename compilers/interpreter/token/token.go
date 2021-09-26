package token

type Type string

const (
	EOF     Type = "EOF"
	Illegal      = "ILLEGAL"

	Semicolon Type = ";"
	Coma           = ","

	// keywords
	Let    Type = "let"
	True        = "true"
	False       = "false"
	If          = "if"
	Else        = "else"
	Return      = "return"

	Identifier = "Identifier"

	// operators
	Plus     Type = "+"
	Minus         = "-"
	Asterisk      = "*"
	Slash         = "/"

	GT = ">"
	GE = ">="
	LT = "<"
	LE = "<="

	Assign Type = "="

	LParen Type = "("
	RParen      = ")"
	LBrace      = "{"
	RBrace      = "}"

	// logical
	Bang Type = "!"

	// datatypes
	Integer Type = "integer"
	String       = "string"

	Function Type = "fn"
)

type Token struct {
	Type  Type
	Value string
}
