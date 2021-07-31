package genetic

import "github.com/MonsieurTa/krpsim/internal/entity"

type Individual struct {
	Genes *GeneList
}

func NewRandomIndividual(size int, p []*entity.Process) *Individual {
	rv := Individual{
		Genes: &GeneList{},
	}
	for i := 0; i < size; i++ {
		rv.Genes.Push(NewRandomGene(p))
	}
	return &rv
}

func (i *Individual) Fitness() int {
	return 0
}
