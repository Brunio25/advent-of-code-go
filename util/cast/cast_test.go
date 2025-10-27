package cast_test

import (
	"advent-of-code-go/util/cast"
	"fmt"
	"testing"
)

var testCases = []struct {
	input    string
	expected int
}{
	{"1", 1},
	{"123", 123},
	{"9901293", 9901293},
	{"-15", -15},
}

func TestToIntSuccess(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`"%s"->%v`, tc.input, tc.expected), func(t *testing.T) {
			if actual := cast.ToInt(tc.input); actual != tc.expected {
				t.Errorf("cast.ToInt() = %v, Expected %v", actual, tc.expected)
			}
		})
	}
}

func TestToIntInvalidInput(t *testing.T) {
	t.Helper()
	defer func() {
		expectedError := `error converting string to int strconv.Atoi: parsing "10fail": invalid syntax`
		if r := recover(); r == nil {
			t.Error("expected panic on cast.ToInt(), but no panic occurred")
		} else if r != expectedError {
			t.Errorf(`expected panic with error "%s" but got "%s"`, expectedError, r)
		}
	}()
	cast.ToInt("10fail")
}
