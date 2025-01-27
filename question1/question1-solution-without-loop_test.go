package question1_test

import (
	"github.com/thaksananan-01/7-solution-test/question1"
	"testing"
)

func TestQuestion1WithoutLoop(t *testing.T) {
	triangle1 := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	expected1 := 237
	result1 := question1.TriangleSumValueWithoutLoop(triangle1)
	if result1 != expected1 {
		t.Errorf("Expected %d, but got %d for triangle1", expected1, result1)
	}

}
