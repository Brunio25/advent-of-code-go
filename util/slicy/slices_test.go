package slicy_test

import (
	"advent-of-code-go/util/slicy"
	"fmt"
	"testing"
)

func TestCountFunc(t *testing.T) {
	testCases := []struct {
		input     []int
		predicate func(int) bool
		expected  int
	}{
		{
			input: []int{1, 2, 3},
			predicate: func(e int) bool {
				return e == 1
			},
			expected: 1,
		},
		{
			input: []int{100, 1000, 10, 1010011, 541234},
			predicate: func(e int) bool {
				return e%10 == 0
			},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`"%v"->%v`, tc.input, tc.expected), func(t *testing.T) {
			if actual := slicy.CountFunc(tc.input, tc.predicate); actual != tc.expected {
				t.Errorf("slicy.CountFunc() = %v, Expected %v", actual, tc.expected)
			}
		})
	}
}
