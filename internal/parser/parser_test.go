package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	dirPath := "../../asset/resources/"
	files := []string{
		"ikea",
		"inception",
		"pomme",
		"recre",
		"simple",
		"steak",
	}

	for _, f := range files {
		filepath := dirPath + f
		p := New()

		_, err := p.Parse(filepath)
		if err != nil {
			t.Fatal(err)
		}
	}
}
