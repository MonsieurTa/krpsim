package simulator

import (
	"fmt"
	"testing"

	"github.com/MonsieurTa/krpsim/internal/parser"
)

func TestIkeaSimulation(t *testing.T) {
	filepath := "../../asset/resources/ikea"
	p := parser.New()

	cfg, err := p.Parse(filepath)
	if err != nil {
		t.Fatal(err)
	}

	populationSize := 100
	genesPerIndividual := 40

	simulator := NewSimulator(&SimulatorConfig{
		KrpSimConfig:               cfg,
		PopulationSize:             populationSize,
		GenesPerIndividual:         genesPerIndividual,
		GenerationLimit:            1000,
		MutationRate:               0.10,
		TournamentPoolSize:         4,
		TournamentSelectionPortion: 0.25,
		ElitismRatio:               0.1,
	})
	simulator.Init()
	result := simulator.Run()[0]

	fitness := simulator.GetFitnesss(result)
	for i := 0; i < int(fitness.Points); i++ {
		fmt.Println(result.Genes()[i].Name)
	}
	fmt.Println("score:", fitness.Score, "points:", fitness.Points, "total delay:", fitness.TotalDelay)
}
