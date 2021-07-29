package parser

import "github.com/MonsieurTa/go-lexer"

type Parser interface {
	Parse(cfg *Config) error
}

type Reader interface {
	NextToken() (lexer.Token, bool)
}
