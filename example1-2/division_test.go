package example1

import (
	"testing"
)

type DivisionTestDefinition struct {
	name     string
	a, b     int
	expected int
	//shouldFail bool
}

var divisionTest []DivisionTestDefinition = []DivisionTestDefinition{
	{"1/1", 1, 1, 1}, //false},
	{"2/3", 2, 3, 0}, //false},
	{"6/2", 6, 2, 3}, //false},
}

func TestAdd_GoodExample(t *testing.T) {
	for _, tt := range divisionTest {
		t.Run(tt.name, func(t *testing.T) {
			result := Division(tt.a, tt.b)
			//result, err := DivisionWithError(tt.a, tt.b)
			//if (err != nil) != tt.shouldFail {
			//	t.Error("unexpected error:", err)
			//}
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
	if result := Division(2, 3); result != 0 {
		t.Errorf("Division(2,3); got %d, want %d", result, 1)
	}
	if result := Division(6, 2); result != 3 {
		t.Errorf("Division(6,2); got %d, want %d", result, 1)
	}
}
