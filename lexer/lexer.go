package lexer

import (
	"github.com/GoyalIshaan/interpreter-in-go/token"
)

type Lexer struct {
	input string
	position int // this is the current character position in the input string
	readPosition int // this is the current reading position i.e next character position in the input string
	ch byte // value of the current character being read
}


// Instantiates a new lexer
func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

// Reads the next character in the input string
func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition++
}

// Peeks the next character in the input string
func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

// Returns the next token in the input string
func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()
	
	switch lexer.ch {
	case '=':
		if lexer.peekChar() == '=' {
			currentChar := lexer.ch
			lexer.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(currentChar) + string(lexer.ch)}
		} else {
			tok = newToken(token.ASSIGN, lexer.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		tok = newToken(token.LPAREN, lexer.ch)
	case ')':
		tok = newToken(token.RPAREN, lexer.ch)
	case ',':
		tok = newToken(token.COMMA, lexer.ch)
	case '+':
		tok = newToken(token.PLUS, lexer.ch)
	case '{':
		tok = newToken(token.LBRACE, lexer.ch)
	case '}':
		tok = newToken(token.RBRACE, lexer.ch)
	case '-':
		tok = newToken(token.MINUS, lexer.ch)
	case '!':
		if lexer.peekChar() == '=' {
			currentChar := lexer.ch
			lexer.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(currentChar) + string(lexer.ch)}
		} else {
			tok = newToken(token.BANG, lexer.ch)
		}
		
	case '*':
		tok = newToken(token.ASTERISK, lexer.ch)
	case '/':
		tok = newToken(token.SLASH, lexer.ch)
	case '<':
		tok = newToken(token.LT, lexer.ch)
	case '>':
		tok = newToken(token.GT, lexer.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookUpIdentifier(tok.Literal)
			return tok
		} else if isDigit(lexer.ch) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()
		 }else {
			tok = newToken(token.ILLEGAL, lexer.ch)
		}
	}
	lexer.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}


func (lexer *Lexer) readIdentifier() string {
	origianlPosition := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[origianlPosition:lexer.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lexer *Lexer) readNumber() string {
	origianlPosition := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[origianlPosition:lexer.position]
}