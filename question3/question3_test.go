package question3_test

import (
	"github.com/thaksananan-01/7-solution-test/question3"
	"testing"
)

func TestQuestion3(t *testing.T) {
	text1 := "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone."
	expected1 := map[string]interface{}{
		"t-bone":   4,
		"fatback":  1,
		"pastrami": 1,
		"pork":     1,
		"meatloaf": 1,
		"jowl":     1,
		"enim":     1,
		"bresaola": 1,
	}
	result1 := question3.PieFireDire(text1)
	if !mapsAreEqual(result1, expected1) {
		t.Errorf("Expected %+v, but got %+v for text1", expected1, result1)
	}

}

// Helper function to compare two maps
func mapsAreEqual(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valueA := range a {
		if valueB, ok := b[key]; !ok || valueA != valueB {
			return false
		}
	}
	return true
}
