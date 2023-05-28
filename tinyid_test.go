package tinyid

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	// Generate a string of length 10
	s := Generate(10)
	if len(s) != 10 {
		t.Errorf("Generate(10) = %v, want length 10", s)
	}
}
