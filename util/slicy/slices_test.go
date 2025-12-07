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

func TestFoldInt(t *testing.T) {
	testCases := []struct {
		input    []int
		op       func(int, int) int
		initial  int
		expected int
	}{
		{
			input: []int{4, 2, 3},
			op: func(a, b int) int {
				return a * b
			},
			initial:  1,
			expected: 24,
		},
		{
			input: []int{100, 1000, 10, 1010011, 541234},
			op: func(a, b int) int {
				return a + b
			},
			initial:  -1,
			expected: 1552354,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`"%v"->%v`, tc.input, tc.expected), func(t *testing.T) {
			if actual := slicy.Fold(tc.input, tc.initial, tc.op); actual != tc.expected {
				t.Errorf("slicy.Fold() = %v, Expected %v", actual, tc.expected)
			}
		})
	}
}

func TestFoldString(t *testing.T) {
	input := []rune{'H', 'E', 'L', 'L', 'O'}
	expected := "HELLO"
	op := func(acc string, v rune) string {
		return acc + string(v)
	}

	t.Run(fmt.Sprintf(`"%v"->%v`, input, expected), func(t *testing.T) {
		if actual := slicy.Fold(input, "", op); actual != expected {
			t.Errorf("slicy.Fold() = %v, Expected %v", actual, expected)
		}
	})
}

func TestReduce(t *testing.T) {
	input := []string{"H", "E", "L", "L", "O"}
	expected := "HELLO"
	op := func(acc string, v string) string {
		return acc + v
	}

	t.Run(fmt.Sprintf(`"%v"->%v`, input, expected), func(t *testing.T) {
		if actual := slicy.Reduce(input, op); actual != expected {
			t.Errorf("slicy.Fold() = %v, Expected %v", actual, expected)
		}
	})
}
