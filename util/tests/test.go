package tests

import (
	"path"
	"reflect"
	"runtime"
	"testing"
)

type TestInputs []struct {
	Name     string
	Input    string
	Expected int
}

func RunWithTestCases(testCases TestInputs, fun func(string) int, t *testing.T) {
	t.Helper()
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if actual := fun(tc.Input); actual != tc.Expected {
				funName := runtime.FuncForPC(reflect.ValueOf(fun).Pointer()).Name()
				t.Errorf("%s() = %v, Expected %v", path.Base(funName), actual, tc.Expected)
			}
		})
	}
}
