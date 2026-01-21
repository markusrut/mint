package lexer

import (
	"mint/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         rune
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lexer.char {
	case '=':
		tok = token.NewToken(token.ASSIGN, lexer.char)
	case '+':
		tok = token.NewToken(token.PLUS, lexer.char)
	case ',':
		tok = token.NewToken(token.COMMA, lexer.char)
	case ';':
		tok = token.NewToken(token.SEMICOLON, lexer.char)
	case '(':
		tok = token.NewToken(token.LPAREN, lexer.char)
	case ')':
		tok = token.NewToken(token.RPAREN, lexer.char)
	case '{':
		tok = token.NewToken(token.LBRACE, lexer.char)
	case '}':
		tok = token.NewToken(token.RBRACE, lexer.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	lexer.readChar()
	return tok
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = []rune(lexer.input)[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}
