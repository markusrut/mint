package parser

import (
	"mint/ast"
	"mint/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram returned nul")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Incorrect statements length, expected 3, got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, test := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, test.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, statement ast.Statement, expectedIdentifier string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("Incorrect TokenLiteral, expected 'let', got %q", statement.TokenLiteral())
	}

	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("Incorrect statement type, expected LetStatement, got %T", statement)
	}

	if letStatement.Name.Value != expectedIdentifier {
		t.Errorf("Incorrect letStatement.Name.Value, expected '%s', got '%s'", expectedIdentifier, letStatement.Name.Value)
	}

	if letStatement.Name.TokenLiteral() != expectedIdentifier {
		t.Errorf("Incorrect letStatement.Name.TokenLiteral(), expected '%s', got '%s'", expectedIdentifier, letStatement.Name.TokenLiteral())
	}

	return true
}
