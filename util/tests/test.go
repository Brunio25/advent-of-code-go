package tests

import (
	"path"
	"reflect"
	"runtime"
	"testing"
)

type TestInputs[T comparable] []struct {
	Name     string
	Input    string
	Expected T
}

func RunWithTestCases[T comparable](testCases TestInputs[T], fun func(string) T, t *testing.T) {
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
