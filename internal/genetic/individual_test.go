package genetic

import (
	"math/rand"
	"testing"
	"time"

	"github.com/MonsieurTa/krpsim/internal/parser"
)

func TestOneRandomIndividual(t *testing.T) {
	filepath := "../../asset/resources/ikea"
	p := parser.New()

	cfg, err := p.Parse(filepath)
	if err != nil {
		t.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())

	expectedSize := 40
	individual := NewRandomIndividual(40, cfg.Processes)
	if expectedSize != individual.TotalGenes() {
		t.Fatalf("expected %d genes, got %d", expectedSize, individual.TotalGenes())
	}
}
