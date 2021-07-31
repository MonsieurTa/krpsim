package genetic

import (
	"math/rand"

	"github.com/MonsieurTa/krpsim/internal/entity"
)

type Gene struct {
	Process *entity.Process
	Next    *Gene
}

func NewGene(p *entity.Process) *Gene {
	return &Gene{p, nil}
}

func NewRandomGene(p []*entity.Process) *Gene {
	idx := rand.Int() % len(p)
	return NewGene(p[idx])
}
