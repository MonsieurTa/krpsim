package simulator

import (
	"math/rand"
	"time"

	"github.com/MonsieurTa/krpsim/internal/genetic"
	"github.com/MonsieurTa/krpsim/internal/utils"
)

type Tournament struct {
	cfg *TournamentConfig
}

type TournamentConfig struct {
	PoolSize int
	Portion  float64
}

// portion is a percentage of the population
func NewTournament(cfg *TournamentConfig) Tournament {
	// TODO: validate cfg
	return Tournament{cfg}
}

func (t *Tournament) Run(fitnesses Fitnesses) []*genetic.Individual {
	maxSize := int(float64(len(fitnesses)) * t.cfg.Portion)

	selection := make([]*genetic.Individual, 0, maxSize)
	for len(selection) < maxSize && len(fitnesses) > 0 {
		i, selected := t.run(fitnesses)

		fitnesses[i] = fitnesses[len(fitnesses)-1]
		fitnesses = fitnesses[:len(fitnesses)-1]

		selection = append(selection, selected.Individual)
	}
	return selection
}

func (t *Tournament) run(fitnesses Fitnesses) (int, *Fitness) {
	rand.Seed(time.Now().UnixNano())

	var best *Fitness
	var randIdx, idx int
	var v *Fitness
	for i := 0; i < t.cfg.PoolSize; i++ {
		for v == nil || v == best {
			randIdx = utils.RandBetween(0, len(fitnesses))
			v = fitnesses[randIdx]
		}
		if (best == nil) ||
			((v.Score > best.Score) ||
				(v.Score == best.Score && v.Points > best.Points)) {
			best, idx = v, randIdx
		}
		v = nil
	}
	return idx, best
}
