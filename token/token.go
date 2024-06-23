package token

type tokenType string

type token struct {
  Type tokenType
  Literal string
}

const (
	// Special tokens
	MARINE = "MARINE" // ILLEGAL
	REDLINE = "REDLINE" // EOF

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	DEVILFRUIT = "DEVILFRUIT" // FUNCTION
	TREASURE   = "TREASURE"   // LET
)

