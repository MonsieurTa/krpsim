package genetic

import (
	"math/rand"
	"testing"
	"time"

	"github.com/MonsieurTa/krpsim/internal/parser"
)

func TestIkeaPopulation(t *testing.T) {
	filepath := "../../asset/resources/ikea"
	p := parser.New()

	cfg, err := p.Parse(filepath)
	if err != nil {
		t.Fatal(err)
	}

	expectedPopulationSize := 40
	expectedGenesPerIndividual := 40

	rand.Seed(time.Now().UnixNano())

	pop := NewRandomPopulation(&Config{
		PopulationSize:     expectedPopulationSize,
		GenesPerIndividual: expectedGenesPerIndividual,
		Processes:          cfg.Processes,
	})

	if len(pop) != expectedPopulationSize {
		t.Fatalf("expected %d population size, got %d", expectedPopulationSize, len(pop))
	}

	for _, v := range pop {
		if v.TotalGenes() != expectedGenesPerIndividual {
			t.Fatalf("expected %d genes, got %d", expectedGenesPerIndividual, v.TotalGenes())
		}
	}
}
