package token

const (
	ILLEGAL TokenType = iota
	EOF

	IDENT
	INT
	STRING

	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH

	EQ
	NOT_EQ

	AND
	OR

	LT
	GT

	COMMA
	COLON
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET

	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

var tokenFriendlyNames = map[TokenType]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	IDENT:  "IDENT",
	INT:    "INT",
	STRING: "STRING",

	ASSIGN:   "=",
	PLUS:     "+",
	MINUS:    "-",
	BANG:     "!",
	ASTERISK: "*",
	SLASH:    "/",

	EQ:     "==",
	NOT_EQ: "!=",

	AND: "&&",
	OR:  "||",

	LT: "<",
	GT: ">",

	COMMA:     ",",
	COLON:     ":",
	SEMICOLON: ";",

	LPAREN:   "(",
	RPAREN:   ")",
	LBRACE:   "{",
	RBRACE:   "}",
	LBRACKET: "[",
	RBRACKET: "]",

	FUNCTION: "fn",
	LET:      "let",
	TRUE:     "true",
	FALSE:    "false",
	IF:       "if",
	ELSE:     "else",
	RETURN:   "return",
}

type TokenType byte

func (tt *TokenType) String() string {
	return tokenFriendlyNames[*tt]
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
