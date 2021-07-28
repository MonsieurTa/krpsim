package lexerstate

import "github.com/MonsieurTa/go-lexer"

const (
	IntToken lexer.TokenType = iota + 1
	IdentToken
	ColonToken
	SemicolonToken
	LPar
	RPar
	EOF
)

var TOKENS = []string{
	lexer.ErrorToken: "ErrorToken",
	IntToken:         "IntToken",
	IdentToken:       "IdentToken",
	ColonToken:       "ColonToken",
	SemicolonToken:   "SemicolonToken",
	LPar:             "LPar",
	RPar:             "RPar",
	EOF:              "EOF",
}

func ToString(t lexer.TokenType) string {
	return TOKENS[t]
}
