package lexer

import (
	"github.com/senicko/writing-an-interpreter-in-go/token"
)

type Lexer struct {
	input        string // source code
	position     int    // current position (points to ch)
	readPosition int    // reading position (points to char after ch)

	// Current character. ch==0 means that we either
	// haven't started lexing yet or we have reached EOF.
	//
	// TODO: Support unicode with runes
	ch byte
}

// New creates a new Lexer instance.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar returns the next character and advances position
// in the input. If after calling readChar() ch==0 it means
// that we've reached EOF
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LBRACE, l.ch)
	case ')':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

// newToken is a util for creating tokens from chars.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
