package parser

import "testing"

func TestGreaterThan(t *testing.T) {
	s1 := scalar{
		scalarType: scalarText,
		text:       "2",
	}
	s2 := scalar{
		scalarType: scalarNumber,
		number:     1,
	}
	if s1.greaterThan(s2) {
		t.Errorf("GT operator unexpectedly returned true for type mismatch")
	}
}
