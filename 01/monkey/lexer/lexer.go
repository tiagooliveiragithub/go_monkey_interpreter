package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New creates a new lexer instance and initializes it with the input string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	// Initialize the lexer by reading the first character
	l.readChar()
	return l
}

// readChar reads the next character in the input string and advances the position pointers
func (l *Lexer) readChar() {
	// Check if we've reached the end of the input string
	if l.readPosition >= len(l.input) {
		l.ch = 0
		// Otherwise, read the next character
	} else {
		l.ch = l.input[l.readPosition]
	}
	// Update the position pointers
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken reads the current character under
// examination and returns a token depending on the character
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Skip whitespaces
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
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
	default:
		// Check if the character is a letter
		if isLetter(l.ch) {
			// Read the identifier/keyword
			tok.Literal = l.readIdentifier()
			// Check if the identifier is a keyword
			tok.Type = token.LookupIdent(tok.Literal)

			// We had to return early here because readIdentifier()
			// advances the position with readChar() and we don't want to do that twice
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			// Read the number
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// newToken creates a new token with the given token type and character
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isLetter checks if the given character is a letter
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// skipWhitespace skips any whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// readNumber reads a number from the input string
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// isDigit checks if the given character is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
