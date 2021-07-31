package genetic

import (
	"io/ioutil"
	"testing"

	"github.com/MonsieurTa/go-lexer"
	"github.com/MonsieurTa/krpsim/internal/entity"
	lexerstate "github.com/MonsieurTa/krpsim/internal/lexer-state"
	"github.com/MonsieurTa/krpsim/internal/parser"
)

func TestIkeaPopulation(t *testing.T) {
	filepath := "../../asset/resources/ikea"
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Error(err)
	}
	l := lexer.New("lexer", string(b), lexerstate.IdentState)
	l.Start()
	p := parser.New(l)

	var cfg entity.Config

	err = p.Parse(&cfg)
	if err != nil {
		t.Fatal(err)
	}

	expectedPopulationSize := 40
	expectedGenesPerIndividual := 40
	pop := NewPopulation(&Config{
		PopulationSize:     expectedPopulationSize,
		GenesPerIndividual: expectedGenesPerIndividual,
		Processes:          cfg.Processes,
	})
	if len(pop.Individuals) != expectedPopulationSize {
		t.Fatalf("expected %d population size, got %d", expectedPopulationSize, len(pop.Individuals))
	}
	for _, v := range pop.Individuals {
		if v.Genes.Size != expectedGenesPerIndividual {
			t.Fatalf("expected %d genes, got %d", expectedGenesPerIndividual, v.Genes.Size)
		}
	}
}
