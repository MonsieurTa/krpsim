package lexerstate

import (
	"io/ioutil"
	"testing"

	"github.com/MonsieurTa/go-lexer"
)

func TestLexer(t *testing.T) {
	dirPath := "../../asset/resources/"
	files := []string{
		"ikea",
		"inception",
		"pomme",
		"recre",
		"simple",
		"steak",
	}

	for _, f := range files {
		filepath := dirPath + f
		b, err := ioutil.ReadFile(filepath)
		if err != nil {
			t.Error(err)
		}
		l := lexer.New("test lexer", string(b), IdentState)
		l.Start()
		rv := []lexer.Token{}
		for {
			if token, done := l.NextToken(); !done {
				rv = append(rv, token)
			} else {
				break
			}
		}
		if len(rv) == 0 {
			t.Error("expected len to be more than 0")
		}
		last := rv[len(rv)-1]
		if last.Type() == lexer.ErrorToken {
			t.Errorf("unexpected error:\n\tat: %d\n\ttype: %s\n\tvalue: %s", last.At(), ToString(last.Type()), last.Value())
		}
	}
}
