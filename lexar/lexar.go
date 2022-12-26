package lexar

import "github.com/kfess/go_interpreter/token"

type Lexar struct {
	input        string // input string
	position     int    // position
	readPosition int    // always point to the next character
	ch           byte   // the character investigating now
}

func isLetter(ch byte) bool {
	if 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' {
		return true
	}
	return false
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func NewToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func New(input string) *Lexar {
	l := &Lexar{input: input}
	l.readChar()

	return l
}

func (l *Lexar) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // Ascii code - NULL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexar) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexar) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexar) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexar) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexar) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = NewToken(token.ASSIGN, '=')
	case ';':
		tok = NewToken(token.SEMICOLON, ';')
	case '(':
		tok = NewToken(token.LPAREN, '(')
	case ')':
		tok = NewToken(token.RPAREN, ')')
	case ',':
		tok = NewToken(token.COMMA, ',')
	case '+':
		tok = NewToken(token.PLUS, '+')
	case '-':
		tok = NewToken(token.MINUS, '-')
	case '!':
		tok = NewToken(token.BANG, '!')
	case '/':
		tok = NewToken(token.SLASH, '/')
	case '*':
		tok = NewToken(token.ASTERISK, '*')
	case '<':
		tok = NewToken(token.LT, '<')
	case '>':
		tok = NewToken(token.GT, '>')
	case '{':
		tok = NewToken(token.LBRACE, '{')
	case '}':
		tok = NewToken(token.RBRACE, '}')
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = NewToken(token.ILLIGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}
