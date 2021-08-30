package utils

import (
	"math/rand"
)

func RandBetween(min, max int) int {
	return rand.Intn(max-min) + min
}

type UniqueRand map[int]bool

func (u UniqueRand) Intn(n int) int {
	for {
		i := rand.Intn(n)
		if !u[i] {
			u[i] = true
			return i
		}
	}
}
