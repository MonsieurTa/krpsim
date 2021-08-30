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
	genesPerIndividual := 100

	simulator := NewSimulator(&SimulatorConfig{
		KrpSimConfig:               cfg,
		PopulationSize:             populationSize,
		GenesPerIndividual:         genesPerIndividual,
		GenerationLimit:            1000,
		MutationRate:               0.1,
		TournamentPoolSize:         genesPerIndividual / 4,
		TournamentSelectionPortion: 0.25,
		ElitismRatio:               0.1,
		CrossoverPoints:            2,
	})
	simulator.Init()
	result := simulator.Run()[0]

	fitness := simulator.GetFitnesss(result)
	totalDelay := 0
	for i := 0; i < int(fitness.SuccessfullGenes); i++ {
		fmt.Println(totalDelay, result.Genes()[i].Name)
		totalDelay += result.Genes()[i].Delay
	}
	fmt.Println(fitness.Store)
	fmt.Println("score:", fitness.Score, "points:", fitness.Points, "total delay:", fitness.TotalDelay)
}
