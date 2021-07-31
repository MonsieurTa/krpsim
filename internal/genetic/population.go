package genetic

import "github.com/MonsieurTa/krpsim/internal/entity"

type Population struct {
	Generation  int
	Individuals []*Individual
}

type Config struct {
	PopulatonSize      int
	GenesPerIndividual int
	Processes          []*entity.Process
}

func NewPopulation(cfg *Config) *Population {
	rv := Population{}

	rv.Individuals = make([]*Individual, cfg.PopulatonSize)
	for i := 0; i < cfg.PopulatonSize; i++ {
		rv.Individuals[i] = NewRandomIndividual(cfg.GenesPerIndividual, cfg.Processes)
	}
	return &rv
}
