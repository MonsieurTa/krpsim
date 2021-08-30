package simulator

import (
	"strconv"

	"github.com/MonsieurTa/krpsim/internal/genetic"
	"github.com/MonsieurTa/krpsim/internal/store"
)

type Fitness struct {
	Individual       *genetic.Individual
	Store            store.Store
	Points           float64
	Score            float64
	TotalDelay       float64
	SuccessfullGenes int
}

type Fitnesses []*Fitness

func (fs Fitnesses) Individuals() []*genetic.Individual {
	rv := make([]*genetic.Individual, len(fs))
	for i, v := range fs {
		rv[i] = v.Individual
	}
	return rv
}

func (f Fitness) String() string {
	pointsStr := strconv.FormatFloat(f.Points, 'f', 2, 64)
	scoreStr := strconv.FormatFloat(f.Score, 'f', 2, 64)
	totalDelayStr := strconv.FormatFloat(f.TotalDelay, 'f', 2, 64)
	totalGenes := f.Individual.TotalGenes()
	firstGeneStr := f.Individual.Genes()[0].Name
	lastGeneStr := f.Individual.Genes()[totalGenes-1].Name
	return "points: " + pointsStr + "\n" +
		" score: " + scoreStr + "\n" +
		" total delay: " + totalDelayStr + "\n" +
		" total genes: " + strconv.FormatInt(int64(totalGenes), 10) + "\n" +
		"\tfirst gene: " + firstGeneStr + "\n" +
		"\tlast gene: " + lastGeneStr + "\n" +
		"\n"
}
