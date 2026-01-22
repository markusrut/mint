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

	for isWhitespace(lexer.char) {
		lexer.readChar()
	}

	switch lexer.char {
	case '=':
		tok = newToken(token.ASSIGN, lexer.char)
	case '+':
		tok = newToken(token.PLUS, lexer.char)
	case '-':
		tok = newToken(token.MINUS, lexer.char)
	case '/':
		tok = newToken(token.SLASH, lexer.char)
	case '*':
		tok = newToken(token.ASTERISK, lexer.char)
	case '!':
		tok = newToken(token.BANG, lexer.char)
	case '<':
		tok = newToken(token.LESS, lexer.char)
	case '>':
		tok = newToken(token.GREATER, lexer.char)
	case ',':
		tok = newToken(token.COMMA, lexer.char)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.char)
	case '(':
		tok = newToken(token.LPAREN, lexer.char)
	case ')':
		tok = newToken(token.RPAREN, lexer.char)
	case '{':
		tok = newToken(token.LBRACE, lexer.char)
	case '}':
		tok = newToken(token.RBRACE, lexer.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.char) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isNumber(lexer.char) {
			tok.Literal = lexer.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.char)
		}
	}

	lexer.readChar()
	return tok
}

func newToken(tokenType token.TokenType, char rune) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
	}
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

func (lexer *Lexer) readIdentifier() string {
	startPosition := lexer.position

	for isLetter(lexer.char) {
		lexer.readChar()
	}

	return lexer.input[startPosition:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	startPosition := lexer.position

	for isNumber(lexer.char) {
		lexer.readChar()
	}

	return lexer.input[startPosition:lexer.position]
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t' || r == '\r'
}

func isLetter(r rune) bool {
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '_'
}

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}
