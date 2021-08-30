package simulator

import (
	"sort"

	"github.com/MonsieurTa/krpsim/internal/entity"
	"github.com/MonsieurTa/krpsim/internal/genetic"
	"github.com/MonsieurTa/krpsim/internal/store"
)

type Simulator struct {
	cfg          *SimulatorConfig
	initialStore store.Store
	population   genetic.Population
	generation   int
}

type SimulatorConfig struct {
	KrpSimConfig               *entity.Config
	PopulationSize             int
	GenesPerIndividual         int
	GenerationLimit            int
	MutationRate               float64
	TournamentPoolSize         int
	TournamentSelectionPortion float64
	ElitismRatio               float64
}

func NewSimulator(cfg *SimulatorConfig) Simulator {
	initialStore := store.New(cfg.KrpSimConfig)

	return Simulator{
		cfg:          cfg,
		initialStore: initialStore,
		population:   nil,
		generation:   0,
	}
}

func (s *Simulator) Init() {
	s.population = genetic.NewRandomPopulation(&genetic.Config{
		PopulationSize:     s.cfg.PopulationSize,
		GenesPerIndividual: s.cfg.GenesPerIndividual,
		Processes:          s.cfg.KrpSimConfig.Processes,
	})
}

func (s *Simulator) Run() genetic.Population {
	var elites []*genetic.Individual

	for i := 0; i < s.cfg.GenerationLimit; i++ {
		fitnesses := s.GetFitnesses()
		elites, fitnesses = s.pullElites(fitnesses)

		tournament := NewTournament(&TournamentConfig{
			PoolSize: s.cfg.TournamentPoolSize,
			Portion:  s.cfg.TournamentSelectionPortion,
		})

		selection := tournament.Run(fitnesses)

		operator := NewGeneticOperator(&OperatorConfig{
			PopulationSize:     s.cfg.PopulationSize - len(elites),
			GenesPerIndividual: s.cfg.GenesPerIndividual,
			MutationRate:       s.cfg.MutationRate,
			Selection:          append(selection, elites[:]...),
			BaseMutations:      s.cfg.KrpSimConfig.Processes,
		})

		nextPopulation := operator.Breed()

		s.population = append(nextPopulation, elites...)

		s.generation++
	}
	return s.population
}

func (s *Simulator) GetFitnesses() Fitnesses {
	rv := make(Fitnesses, len(s.population))

	for i, individual := range s.population {
		rv[i] = s.GetFitnesss(individual)
	}

	sort.SliceStable(rv, func(i, j int) bool {
		return ((rv[i].Score > rv[j].Score) ||
			(rv[i].Score == rv[j].Score && rv[i].Points > rv[j].Points))
	})
	return rv
}

func (s *Simulator) GetFitnesss(individual *genetic.Individual) *Fitness {
	store := s.initialStore.Duplicate()
	rv := Fitness{Individual: individual}

	for _, v := range individual.Genes() {
		if store.ConsumeIfAvailable(v.Needs) {
			rv.Points++
			rv.TotalDelay += float64(v.Delay)
			for _, goal := range s.cfg.KrpSimConfig.Goals {
				_, ok := v.Results[goal]
				if ok {
					rv.Score = rv.Points
					if s.cfg.KrpSimConfig.OptimizeTime && rv.TotalDelay != 0 {
						rv.Score /= rv.TotalDelay
					}
				}
			}
			store.BatchStore(v.Results)
		} else {
			break
		}
	}
	return &rv
}

func (s *Simulator) pullElites(fitnesses Fitnesses) ([]*genetic.Individual, Fitnesses) {
	maxSize := int(s.cfg.ElitismRatio * float64(s.cfg.PopulationSize))
	rv := make([]*genetic.Individual, maxSize)

	for i := range rv {
		rv[i] = fitnesses[i].Individual
	}
	return rv, fitnesses[maxSize:]
}
