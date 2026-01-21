package lexer

import (
	"mint/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, testToken := range tests {
		token := l.NextToken()

		if token.Type != testToken.expectedType {
			t.Fatalf("tests[%d] - Type is incorrect. Expected %q but got %q", i, testToken.expectedType, token.Type)
		}

		if token.Literal != testToken.expectedLiteral {
			t.Fatalf("tests[%d] - Literal is incorrect. Expected %q but got %q", i, testToken.expectedLiteral, token.Literal)
		}
	}
}
