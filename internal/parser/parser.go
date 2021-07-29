package parser

import (
	"fmt"
	"strconv"

	"github.com/MonsieurTa/go-lexer"
	lexerstate "github.com/MonsieurTa/krpsim/internal/lexer-state"
)

type parser struct {
	r      Reader
	tokens *Stack
	cfg    *Config
}

func New(r Reader) Parser {
	if r == nil {
		panic("parser: got nil reader")
	}
	return &parser{
		r:      r,
		tokens: &Stack{},
	}
}

func (p *parser) Parse(cfg *Config) error {
	p.cfg = cfg
	p.pullTokens()
	return p.parse()
}

func (p *parser) parse() error {
	var err error
	p.cfg.Stocks, err = p.parseStocks()
	if err != nil {
		return err
	}
	p.cfg.Processes, err = p.parseProcesses()
	if err != nil {
		return err
	}
	return nil
}

func (p *parser) pullTokens() {
	for {
		if t, done := p.r.NextToken(); !done {
			p.tokens.PushBack(t)
		} else {
			break
		}
	}
}

func (p *parser) parseStocks() ([]*Stock, error) {
	rv := []*Stock{}
	for {
		s, err := p.parseStock()
		if err != nil {
			return nil, err
		}
		if s == nil {
			return rv, nil
		}
		rv = append(rv, s)

		p.tokens.IgnoreIf([]lexer.TokenType{lexerstate.SemicolonToken})
	}
}

func (p *parser) parseStock() (*Stock, error) {
	keyNode := p.tokens.PopFront()
	sepNode := p.tokens.PopFront()
	valueNode := p.tokens.PopFront()

	if keyNode == nil || sepNode == nil || valueNode == nil {
		return nil, fmt.Errorf("unexpected nil, keyNode=%v, sepNode=%v, valueNode=%v", keyNode, sepNode, valueNode)
	}
	if !isStock(keyNode, sepNode, valueNode) {
		p.tokens.BatchPushFront([]*StackNode{keyNode, sepNode, valueNode})
		return nil, nil
	}
	v, err := strconv.Atoi(valueNode.Val.Value())
	if err != nil {
		return nil, err
	}
	return &Stock{keyNode.Val.Value(), v}, nil
}

func isStock(key, sep, value *StackNode) bool {
	return key.IsType(lexerstate.IdentToken) &&
		sep.IsType(lexerstate.ColonToken) &&
		value.IsType(lexerstate.IntToken)
}

func (p *parser) parseProcesses() ([]*Process, error) {
	rv := []*Process{}
	for {
		s, err := p.parseProcess()
		if s == nil && err == nil {
			return rv, nil
		}
		if err != nil {
			return nil, err
		}
		rv = append(rv, s)
	}
}

func (p *parser) parseProcess() (*Process, error) {
	process := p.tokens.PopFront()

	if !process.IsType(lexerstate.IdentToken) {
		return nil, fmt.Errorf("expected ident token, got %v", lexerstate.ToString(process.Val.Type()))
	}
	p.tokens.IgnoreIf([]lexer.TokenType{lexerstate.ColonToken, lexerstate.LPar})

	needs, err := p.parseStocks()
	if err != nil || len(needs) == 0 {
		return nil, nil
	}

	p.tokens.IgnoreIf([]lexer.TokenType{lexerstate.ColonToken, lexerstate.LPar, lexerstate.RPar})

	results, err := p.parseStocks()
	if err != nil {
		return nil, err
	}

	p.tokens.IgnoreIf([]lexer.TokenType{lexerstate.ColonToken, lexerstate.RPar})

	delayNode := p.tokens.PopFront()
	if delayNode == nil {
		return nil, fmt.Errorf("expected delay node, got %v", delayNode)
	}
	if !delayNode.IsType(lexerstate.IntToken) {
		return nil, fmt.Errorf("expected int token, got %v", lexerstate.ToString(delayNode.Val.Type()))
	}
	delay, err := strconv.Atoi(delayNode.Val.Value())
	if err != nil {
		return nil, err
	}
	return &Process{
		Name:    process.Val.Value(),
		Needs:   needs,
		Results: results,
		Delay:   delay,
	}, nil
}

func (p *parser) parseOptimize() {

}
