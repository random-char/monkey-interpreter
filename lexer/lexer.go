package lexer

import "monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int  //current position in input
	readPosition int  //current reading position
	ch           byte //current char
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()

	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = newToken(token.EQ, "==")
		} else {
			tok = newToken(token.ASSIGN, string(l.ch))
		}
	case '+':
		tok = newToken(token.PLUS, string(l.ch))
	case '-':
		tok = newToken(token.MINUS, string(l.ch))
	case '*':
		tok = newToken(token.ASTERISK, string(l.ch))
	case '/':
		tok = newToken(token.SLASH, string(l.ch))
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = newToken(token.NOT_EQ, "!=")
		} else {
			tok = newToken(token.BANG, string(l.ch))
		}
	case '&':
		if l.peekChar() == '&' {
			l.readChar()
			tok = newToken(token.AND, "&&")
		}
	case '|':
		if l.peekChar() == '|' {
			l.readChar()
			tok = newToken(token.OR, "||")
		}
	case '<':
		tok = newToken(token.LT, string(l.ch))
	case '>':
		tok = newToken(token.GT, string(l.ch))
	case ':':
		tok = newToken(token.COLON, string(l.ch))
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch))
	case '(':
		tok = newToken(token.LPAREN, string(l.ch))
	case ')':
		tok = newToken(token.RPAREN, string(l.ch))
	case '{':
		tok = newToken(token.LBRACE, string(l.ch))
	case '}':
		tok = newToken(token.RBRACE, string(l.ch))
	case '[':
		tok = newToken(token.LBRACKET, string(l.ch))
	case ']':
		tok = newToken(token.RBRACKET, string(l.ch))
	case ',':
		tok = newToken(token.COMMA, string(l.ch))
	case '"':
		tok = newToken(token.STRING, l.readString())
	case 0:
		tok = newToken(token.EOF, "")
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			return newToken(token.INT, l.readNumber())
		} else {
			tok = newToken(token.ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// todo
// add unterminated string error
// add support for \"
func (l *Lexer) readString() string {
	position := l.position + 1

	for {
		l.readChar()

		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
