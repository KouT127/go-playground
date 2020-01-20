package lexer

import (
	"go-playground/token"
	"testing"
)

//　テスト
func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x,y){
		x + y;
	};
	let result = add(five,ten)`

	tests := []struct {
		expectedType    token.TypeToken
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.INDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
	}
	l := createLexer(input)
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
