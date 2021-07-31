package parser

import (
	"io/ioutil"
	"testing"

	"github.com/MonsieurTa/go-lexer"
	"github.com/MonsieurTa/krpsim/internal/entity"
	lexerstate "github.com/MonsieurTa/krpsim/internal/lexer-state"
)

func TestParser(t *testing.T) {
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
		l := lexer.New("test lexer", string(b), lexerstate.IdentState)
		l.Start()

		p := New(l)

		var cfg entity.Config

		if p.Parse(&cfg) != nil {
			t.Fatal(err)
		}
	}
}
