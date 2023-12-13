package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5
	let ten = 10

	let add = function(x, y) {
		return x + y;
	}

	let result = add(five, ten);

	!-/*5;
	5 < 10 > 5;
`

	expectedTokens := []token.Token{
		{token.LET, "let"},
		{token.ID, "five"},
		{token.ASSIGN, "="},
		{token.NUMBER, "5"},
		{token.LET, "let"},
		{token.ID, "ten"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10"},
		{token.LET, "let"},
		{token.ID, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "function"},
		{token.LPAREN, "("},
		{token.ID, "x"},
		{token.COMMA, ","},
		{token.ID, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.ID, "x"},
		{token.PLUS, "+"},
		{token.ID, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.LET, "let"},
		{token.ID, "result"},
		{token.ASSIGN, "="},
		{token.ID, "add"},
		{token.LPAREN, "("},
		{token.ID, "five"},
		{token.COMMA, ","},
		{token.ID, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.NUMBER, "5"},
		{token.LESS, "<"},
		{token.NUMBER, "10"},
		{token.GREATER, ">"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)
	assertTokens(t, expectedTokens, lexer)
}

func assertTokens(t *testing.T, expectedTokens []token.Token, lexer *Lexer) {
	for i, expected := range expectedTokens {
		actualToken := lexer.NextToken()
		if actualToken.Type != expected.Type || actualToken.Literal != expected.Literal {
			t.Fatalf("case[%d] failed, expected=%s, actual=%s", i, expected, actualToken)
		}
	}
}
