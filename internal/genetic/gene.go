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

type GeneList struct {
	Head *Gene
	Tail *Gene
	Size int
}

func (c *GeneList) Push(g *Gene) {
	if c.Head == nil && c.Tail == nil {
		c.Head = g
		c.Tail = g
	} else {
		c.Tail.Next = g
		c.Tail = g
	}
	c.Tail.Next = nil
	c.Size++
}
