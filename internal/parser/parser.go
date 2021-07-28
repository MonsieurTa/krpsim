package parser

import (
	"sync"

	"github.com/MonsieurTa/go-lexer"
)

type parser struct {
	input      chan lexer.Token
	tokenStack safeTokenStack
}

type safeTokenStack struct {
	mu sync.Mutex
	v  []lexer.Token
}

func (s *safeTokenStack) push(t lexer.Token) {
	defer s.mu.Unlock()
	s.mu.Lock()

	s.v = append(s.v, t)
}

func (s *safeTokenStack) pop() lexer.Token {
	defer s.mu.Unlock()
	s.mu.Lock()

	if len(s.v) == 0 {
		return nil
	}

	var rv lexer.Token
	rv, s.v = s.v[0], s.v[1:]
	return rv
}

func New(c chan lexer.Token) Parser {
	if c == nil {
		panic("parser: got nil channel")
	}
	return &parser{
		input:      c,
		tokenStack: safeTokenStack{v: []lexer.Token{}},
	}
}

func (p *parser) Start() {
	go p.pullTokens()

}

func (p *parser) pullTokens() {
	for v := range p.input {
		p.tokenStack.push(v)
	}
}
