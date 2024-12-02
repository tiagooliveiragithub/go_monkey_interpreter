package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {

	input := `=+(){},;`

	// Define the expected token types and literals
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
		{token.EOF, ""},
	}

	// Create a new lexer instance
	l := New(input)

	// Iterate over the test cases and compare the expected token type and literal
	for i, tt := range tests {
		// Get the next token from the lexer
		tok := l.NextToken()

		// Compare the expected token type with the actual token type
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		// Compare the expected token literal with the actual token literal
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
