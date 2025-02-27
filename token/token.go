package token

type TokenType string

const (
	// Illegacl character we don't recognize
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals

	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Operators

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	EQ       = "="
	NOT_EQ   = "!="
	LT       = "<"
	GT       = ">"
	SLASH    = "/"
	ASTERISK = "*"
	BANG     = "!"

	// Delimiters

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COLON     = ":"

	// Keywords

	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"let":    LET,
	"fn":     FUNCTION,
	"return": RETURN,
}

func LookupIdentifier(s string) TokenType {
	if token, ok := keywords[s]; ok {
		return token
	}
	return IDENT
}
