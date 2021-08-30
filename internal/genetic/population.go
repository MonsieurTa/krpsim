package genetic

import (
	"math/rand"
	"time"

	"github.com/MonsieurTa/krpsim/internal/entity"
)

type Population []*Individual

type Config struct {
	PopulationSize     int
	GenesPerIndividual int
	Processes          []*entity.Process
}

func NewRandomPopulation(cfg *Config) Population {
	rand.Seed(time.Now().UnixNano())

	rv := make(Population, cfg.PopulationSize)
	for i := 0; i < cfg.PopulationSize; i++ {
		rv[i] = NewRandomIndividual(cfg.GenesPerIndividual, cfg.Processes)
	}
	return rv
}

func (p Population) Size() int {
	return len(p)
}
