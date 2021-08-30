package genetic

import (
	"strings"

	"github.com/MonsieurTa/krpsim/internal/entity"
	"github.com/MonsieurTa/krpsim/internal/utils"
)

type Individual struct {
	genes Genes
}

func NewIndividual(size int) Individual {
	return Individual{make(Genes, 0, size)}
}

func NewRandomIndividual(size int, p []*entity.Process) *Individual {
	rv := NewIndividual(size)

	for i := 0; i < size; i++ {
		rv.genes.Push(NewRandomGene(p))
	}
	return &rv
}

func (ind Individual) String() string {
	var builder strings.Builder

	for _, v := range ind.genes {
		builder.WriteString(v.Name + "\n")
	}
	return builder.String()
}

func (ind Individual) Genes() Genes {
	return ind.genes
}

func (ind *Individual) AddGene(v Gene) {
	ind.genes.Push(v)
}

func (ind Individual) TotalGenes() int {
	return ind.genes.Size()
}

func (ind *Individual) Mutate(newGene Gene) {
	randIdx := utils.RandBetween(0, ind.TotalGenes())
	ind.genes[randIdx] = newGene
}
