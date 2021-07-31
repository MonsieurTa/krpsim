package genetic

import "github.com/MonsieurTa/krpsim/internal/entity"

type Individual struct {
	Head *Gene
	Tail *Gene
	Size int
}

func (c *Individual) Push(g *Gene) {
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

func NewRandomIndividual(size int, p []*entity.Process) *Individual {
	rv := Individual{}
	for i := 0; i < size; i++ {
		rv.Push(NewRandomGene(p))
	}
	return &rv
}

func (i *Individual) Fitness() int {
	return 0
}
