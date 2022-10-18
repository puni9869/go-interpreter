package lexer

import (
	"play/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
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

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
		l.readChar()
	case '+':
		tok = newToken(token.PLUS, l.ch)
		l.readChar()
	case '-':
		tok = newToken(token.MINUS, l.ch)
		l.readChar()
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
		l.readChar()
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
		l.readChar()
	case '<':
		tok = newToken(token.LT, l.ch)
		l.readChar()
	case '>':
		tok = newToken(token.GT, l.ch)
		l.readChar()
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
		l.readChar()
	case ',':
		tok = newToken(token.COMMA, l.ch)
		l.readChar()
	case '{':
		tok = newToken(token.LBRACE, l.ch)
		l.readChar()
	case '}':
		tok = newToken(token.RBRACE, l.ch)
		l.readChar()
	case '(':
		tok = newToken(token.LPAREN, l.ch)
		l.readChar()
	case ')':
		tok = newToken(token.RPAREN, l.ch)
		l.readChar()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		l.readChar()
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGLE, l.ch)
		}
	}

	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]

}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
