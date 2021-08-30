package parser

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/MonsieurTa/go-lexer"
	"github.com/MonsieurTa/krpsim/internal/entity"
	lexerstate "github.com/MonsieurTa/krpsim/internal/lexer-state"
)

type parser struct {
	r      Reader
	tokens *Stack
}

func New() Parser {
	return &parser{
		tokens: &Stack{},
	}
}

func (p *parser) Parse(filepath string) (*entity.Config, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	l := lexer.New("lexer", string(b), lexerstate.IdentState)
	p.r = l
	l.Start()

	p.pullTokens()
	return p.parse()
}

func (p *parser) parse() (*entity.Config, error) {
	var cfg entity.Config
	var err error

	cfg.Stocks, err = p.parseStocks()
	if err != nil {
		return nil, err
	}
	cfg.Processes, err = p.parseProcesses()
	if err != nil {
		return nil, err
	}
	cfg.OptimizeTime, cfg.Goals, err = p.parseGoals()
	if err != nil {
		return nil, err
	}
	return &cfg, nil
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

func (p *parser) parseStocks() (map[string]int, error) {
	rv := map[string]int{}
	for {
		key, value, err := p.parseStock()
		if err != nil {
			return nil, err
		}
		if key == "" && value == -1 {
			return rv, nil
		}
		rv[key] = value

		p.tokens.IgnoreIf([]lexer.TokenType{lexerstate.SemicolonToken})
	}
}

func (p *parser) parseStock() (key string, value int, err error) {
	keyNode := p.tokens.PopFront()
	sepNode := p.tokens.PopFront()
	valueNode := p.tokens.PopFront()

	if keyNode == nil || sepNode == nil || valueNode == nil {
		return "", -1, fmt.Errorf("unexpected nil, keyNode=%v, sepNode=%v, valueNode=%v", keyNode, sepNode, valueNode)
	}
	if !isStock(keyNode, sepNode, valueNode) {
		p.tokens.BatchPushFront([]*StackNode{keyNode, sepNode, valueNode})
		return "", -1, nil
	}
	v, err := strconv.Atoi(valueNode.Val.Value())
	if err != nil {
		return "", -1, err
	}
	return keyNode.Val.Value(), v, nil
}

func isStock(key, sep, value *StackNode) bool {
	return key.IsType(lexerstate.IdentToken) &&
		sep.IsType(lexerstate.ColonToken) &&
		value.IsType(lexerstate.IntToken)
}

func (p *parser) parseProcesses() ([]*entity.Process, error) {
	rv := make([]*entity.Process, 0)
	for {
		process, err := p.parseProcess()
		if process == nil && err == nil {
			return rv, nil
		}
		if err != nil {
			return nil, err
		}
		rv = append(rv, process)
	}
}

func (p *parser) parseProcess() (*entity.Process, error) {
	process := p.tokens.PopFront()

	if process == nil {
		return nil, fmt.Errorf("expected ident token, got %v", process)
	} else if !process.IsType(lexerstate.IdentToken) {
		return nil, fmt.Errorf("expected ident token, got %v", lexerstate.ToString(process.Val.Type()))
	} else if process.Val.Value() == "optimize" {
		p.tokens.PushFront(process)
		return nil, nil
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
	return &entity.Process{
		Name:    process.Val.Value(),
		Needs:   needs,
		Results: results,
		Delay:   delay,
	}, nil
}

func (p *parser) parseGoals() (bool, []string, error) {
	optimize := p.tokens.PopFront()

	if optimize == nil {
		return false, nil, fmt.Errorf("expected ident token, got %v", optimize)
	}
	if !optimize.IsType(lexerstate.IdentToken) {
		return false, nil, fmt.Errorf("expected ident token, got %v", lexerstate.ToString(optimize.Val.Type()))
	}
	p.tokens.IgnoreIf([]lexer.TokenType{lexerstate.ColonToken, lexerstate.LPar})

	return getGoalTokens(p.tokens)
}

func getGoalTokens(tokens *Stack) (bool, []string, error) {
	memo := map[string]bool{} // avoid duplicates

	t := tokens.PopFront()
	for t != nil && !t.IsType(lexerstate.RPar) {
		if t.IsType(lexerstate.IdentToken) {
			memo[t.Val.Value()] = true
		} else if !t.IsType(lexerstate.SemicolonToken) {
			return false, nil, fmt.Errorf("unexpected token %s, expected ident token or semicolon token", t.Val.Value())
		}
		t = tokens.PopFront()
	}
	if t == nil {
		return false, nil, fmt.Errorf("unexpected EOF, expected ')'")
	}

	rv := []string{}
	optimizeTime := false
	for name := range memo {
		if name == "time" {
			optimizeTime = true
		} else {
			rv = append(rv, name)
		}
	}
	return optimizeTime, rv, nil
}
