package example1

import (
	"testing"
)

type DivisionTestDefinition struct {
	name     string
	a, b     int
	expected int
}

var divisionTest []DivisionTestDefinition = []DivisionTestDefinition{
	{"1/1", 1, 1, 1},
	{"2/3", 2, 3, 0},
	{"6/2", 6, 2, 3},
}

func TestAdd_GoodExample(t *testing.T) {
	for _, tt := range divisionTest {
		t.Run(tt.name, func(t *testing.T) {
			result := Division(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestAdd_BadExample(t *testing.T) {
	if result := Division(1, 1); result != 1 {
		t.Errorf("Division(1,1); got %d, want %d", result, 1)
	}
}
