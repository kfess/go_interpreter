package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// token 一覧
const (
	ILLIGAL = "ILLIAGL"
	EOF     = "EOF"

	// identifier
	IDENT = "IDENT" // add, x, y, foo, bar
	INT   = "INT"

	// operator
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	LE     = "<="
	GE     = ">="
	EQ     = "=="
	NOT_EQ = "!="

	// delimiter
	COMMA     = ","
	SEMICOLON = ";"

	// Parentheses
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keyword
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
