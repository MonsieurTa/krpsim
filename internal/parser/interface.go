package parser

import (
	"github.com/MonsieurTa/go-lexer"
	"github.com/MonsieurTa/krpsim/internal/entity"
)

type Parser interface {
	Parse(cfg *entity.Config) error
}

type Reader interface {
	NextToken() (lexer.Token, bool)
}
