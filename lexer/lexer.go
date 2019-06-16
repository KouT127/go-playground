package lexer

import (
	"go-playground/token"
	"testing"
)

type Lexer struct {
	input string
	// 入力における現在の位置
	position int
	// これから読み込む位置
	readPosition int
	// 現在調査中の文字
	ch byte
}

//　テスト
func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TypeToken
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
		{token.EOF, "EOF"},
	}
	l := new(input)
	for i, tt := range tests {
		tok := l.nextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

// 次のトークンを読み込む
func (l *Lexer) nextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

//　Charを読み込む
func (l *Lexer) readChar() {
	// 次のポジション　>= インプットの長さ
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

//　新しいTokenを生成
func newToken(typeToken token.TypeToken, ch byte) token.Token {
	return token.Token{Type: typeToken, Literal: string(ch)}
}

// 引数の文字を元に、読みとった後の字句解析の構造体を返す。
func new(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
