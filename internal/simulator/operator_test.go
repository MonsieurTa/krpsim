package simulator

import (
	"testing"

	"github.com/MonsieurTa/krpsim/internal/parser"
)

func TestIkeaGeneticOperator(t *testing.T) {
	filepath := "../../asset/resources/ikea"
	p := parser.New()

	cfg, err := p.Parse(filepath)
	if err != nil {
		t.Fatal(err)
	}

	populationSize := 40
	genesPerIndividual := 40

	simulator := NewSimulator(&SimulatorConfig{
		KrpSimConfig:       cfg,
		PopulationSize:     populationSize,
		GenesPerIndividual: genesPerIndividual,
	})

	simulator.Init()

	poolSize := populationSize / 4
	tournament := NewTournament(&TournamentConfig{
		PoolSize: poolSize,
		Portion:  0.25,
	})
	fitnesses := simulator.GetFitnesses()
	selection := tournament.Run(fitnesses)

	// find a better test
	expectedSize := 10
	if len(selection) != expectedSize {
		t.Fatalf("expected %d selection size, got %d", expectedSize, len(selection))
	}

	operator := NewGeneticOperator(&OperatorConfig{
		PopulationSize:     populationSize,
		GenesPerIndividual: genesPerIndividual,
		MutationRate:       0.10,
		Selection:          selection,
		BaseMutations:      cfg.Processes,
	})

	newInds := operator.Breed()

	if populationSize != len(newInds) {
		t.Fatalf("expected %d, got %d", populationSize, len(newInds))
	}
}
