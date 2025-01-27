package question2_test

import (
	"reflect"
	"github.com/thaksananan-01/7-solution-test/question2"
	"testing"
)

type testCase struct {
	encoded   string
	expected  []int
}

func TestQuestion2(t *testing.T) {
	tests := []testCase {
		{encoded: "L", expected: []int{1, 0}},
		{encoded: "R", expected: []int{0, 1}},
		{encoded: "=", expected: []int{0, 0}},
		{encoded: "LL", expected: []int{2, 1, 0}},
		{encoded: "RR", expected: []int{0, 1, 2}},
		{encoded: "L=", expected: []int{1, 0, 0}},
		{encoded: "R=", expected: []int{0, 1, 1}},
	}

	for _, test := range tests {
		t.Run(test.encoded, func(t *testing.T) {
			result := question2.DecodeString(test.encoded)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For encoded=%q, expected %v, but got %v", test.encoded, test.expected, result)
			}
		})
	}
}