package genetic

import (
	"math/rand"

	"github.com/MonsieurTa/krpsim/internal/entity"
)

type Gene *entity.Process
type Genes []Gene

func NewGene(p *entity.Process) Gene {
	return p
}

func NewRandomGene(p []*entity.Process) Gene {
	idx := rand.Int() % len(p)
	return NewGene(p[idx])
}

func (gs *Genes) Push(v Gene) {
	*gs = append(*gs, v)
}

func (gs Genes) Size() int {
	return len(gs)
}
