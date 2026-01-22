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
		if lexer.peekChar() == '=' {
			tok = lexer.makeTwoCharToken(token.EQUALS)
		} else {
			tok = lexer.makeToken(token.ASSIGN)
		}
	case '+':
		tok = lexer.makeToken(token.PLUS)
	case '-':
		tok = lexer.makeToken(token.MINUS)
	case '/':
		tok = lexer.makeToken(token.SLASH)
	case '*':
		tok = lexer.makeToken(token.ASTERISK)
	case '!':
		if lexer.peekChar() == '=' {
			tok = lexer.makeTwoCharToken(token.NOT_EQUALS)
		} else {
			tok = lexer.makeToken(token.BANG)
		}
	case '<':
		tok = lexer.makeToken(token.LESS)
	case '>':
		tok = lexer.makeToken(token.GREATER)
	case ',':
		tok = lexer.makeToken(token.COMMA)
	case ';':
		tok = lexer.makeToken(token.SEMICOLON)
	case '(':
		tok = lexer.makeToken(token.LPAREN)
	case ')':
		tok = lexer.makeToken(token.RPAREN)
	case '{':
		tok = lexer.makeToken(token.LBRACE)
	case '}':
		tok = lexer.makeToken(token.RBRACE)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.char) {
			return lexer.makeIdentifierToken()
		} else if isNumber(lexer.char) {
			return lexer.makeNumberToken()
		} else {
			tok = lexer.makeToken(token.ILLEGAL)
		}
	}

	lexer.readChar()
	return tok
}

func (lexer *Lexer) makeToken(tokenType token.TokenType) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(lexer.char),
	}
}

func (lexer *Lexer) makeTwoCharToken(tokenType token.TokenType) token.Token {
	char := lexer.char
	lexer.readChar()
	return token.Token{
		Type:    tokenType,
		Literal: string(char) + string(lexer.char),
	}
}

func (lexer *Lexer) makeIdentifierToken() token.Token {
	identifier := lexer.readIdentifier()
	return token.Token{
		Type:    token.LookupIdentifier(identifier),
		Literal: identifier,
	}
}

func (lexer *Lexer) makeNumberToken() token.Token {
	number := lexer.readNumber()
	return token.Token{
		Type:    token.INT,
		Literal: number,
	}
}

func (lexer *Lexer) readChar() {
	lexer.char = lexer.peekChar()
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) peekChar() rune {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return []rune(lexer.input)[lexer.readPosition]
	}
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
