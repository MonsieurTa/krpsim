package simulator

import (
	"math/rand"
	"sort"

	"github.com/MonsieurTa/krpsim/internal/entity"
	"github.com/MonsieurTa/krpsim/internal/genetic"
	"github.com/MonsieurTa/krpsim/internal/utils"
)

type GeneticOperator struct {
	cfg *OperatorConfig
}

type OperatorConfig struct {
	PopulationSize     int
	GenesPerIndividual int
	MutationRate       float64
	Selection          []*genetic.Individual
	BaseMutations      []*entity.Process
}

func NewGeneticOperator(cfg *OperatorConfig) *GeneticOperator {
	// TODO: validate cfg
	return &GeneticOperator{cfg}
}

func (o *GeneticOperator) Breed() genetic.Population {
	newPopulation := o.crossover()
	o.mutate(newPopulation)
	return newPopulation
}

func (o *GeneticOperator) crossover() genetic.Population {
	rv := make(genetic.Population, 0, o.cfg.PopulationSize)

	for len(rv) < o.cfg.PopulationSize {
		children := o.breed()
		rv = append(rv, children[:]...)
	}
	return rv
}

func (o *GeneticOperator) breed() [2]*genetic.Individual {
	return o.kPointCrossover(2)
}

func (o *GeneticOperator) onePointCrossover(k int) [2]*genetic.Individual {
	father, mother := o.getRandomParents()

	point := utils.RandBetween(0, o.cfg.GenesPerIndividual)

	firstChild := genetic.NewIndividual(o.cfg.GenesPerIndividual)
	secondChild := genetic.NewIndividual(o.cfg.GenesPerIndividual)

	for i := 0; i < o.cfg.GenesPerIndividual; i++ {
		if i < point {
			firstChild.AddGene(father.Genes()[i])
			secondChild.AddGene(mother.Genes()[i])
		} else {
			firstChild.AddGene(mother.Genes()[i])
			secondChild.AddGene(father.Genes()[i])
		}
	}
	return [2]*genetic.Individual{&firstChild, &secondChild}
}

func (o *GeneticOperator) kPointCrossover(k int) [2]*genetic.Individual {
	uniqueRand := utils.UniqueRand{}
	points := make([]int, k)
	for i := range points {
		points[i] = uniqueRand.Intn(o.cfg.GenesPerIndividual)
	}
	sort.Ints(points)

	father, mother := o.getRandomParents()
	firstChild := genetic.NewIndividual(o.cfg.GenesPerIndividual)
	secondChild := genetic.NewIndividual(o.cfg.GenesPerIndividual)

	for i := 0; i < o.cfg.GenesPerIndividual; i++ {
		if len(points) > 0 && i >= points[0] {
			points = points[1:]
			// swap
			tmp := firstChild
			firstChild = secondChild
			secondChild = tmp
		}
		firstChild.AddGene(father.Genes()[i])
		secondChild.AddGene(mother.Genes()[i])
	}
	return [2]*genetic.Individual{&firstChild, &secondChild}
}

func (o *GeneticOperator) getRandomParents() (*genetic.Individual, *genetic.Individual) {
	fatherIdx := utils.RandBetween(0, len(o.cfg.Selection))
	motherIdx := utils.RandBetween(0, len(o.cfg.Selection))
	for motherIdx == fatherIdx {
		motherIdx = utils.RandBetween(0, len(o.cfg.Selection))
	}
	return o.cfg.Selection[fatherIdx], o.cfg.Selection[motherIdx]
}

func (o *GeneticOperator) mutate(inds []*genetic.Individual) {
	for _, v := range inds {
		canMutate := rand.Float64() <= o.cfg.MutationRate
		if canMutate {
			randGene := genetic.NewRandomGene(o.cfg.BaseMutations)
			v.Mutate(randGene)
		}
	}
}
