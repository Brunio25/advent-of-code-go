package geom_test

import (
	"advent-of-code-go/util/geom"
	"fmt"
	"testing"
)

func TestCoordinates2D_WithX(t *testing.T) {
	original := geom.Coordinates2D{X: 0, Y: 0}
	expected := geom.Coordinates2D{X: 27, Y: 0}
	if actual := original.WithX(27); actual != expected {
		t.Errorf("coordinates.Coordinates2d.WithX() = Coordinates2D%v, Expected Coordinates2D%v", actual, expected)
	}
}

func TestCoordinates2D_WithY(t *testing.T) {
	original := geom.Coordinates2D{X: 0, Y: 0}
	expected := geom.Coordinates2D{X: 0, Y: 27}
	if actual := original.WithY(27); actual != expected {
		t.Errorf("coordinates.Coordinates2d.WithY() = Coordinates2D%v, Expected Coordinates2D%v", actual, expected)
	}
}

func TestCoordinates2D_EuclideanDistance(t *testing.T) {
	testCases := []struct {
		p1       geom.Coordinates2D
		p2       geom.Coordinates2D
		expected float64
	}{
		{
			p1:       geom.Coordinates2D{X: 1, Y: 1},
			p2:       geom.Coordinates2D{X: 10, Y: 10},
			expected: 12.727922061357855,
		},
		{
			p1:       geom.Coordinates2D{X: 3, Y: -1},
			p2:       geom.Coordinates2D{X: 12, Y: 57},
			expected: 58.69412236331676,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`"%v-%v"->%v`, tc.p1, tc.p2, tc.expected), func(t *testing.T) {
			if actual := tc.p1.EuclideanDistance(tc.p2); actual != tc.expected {
				t.Errorf("coordinates.Coordinates2d.EuclidenaDistance() = %v, Expected %v", actual, tc.expected)
			}
		})
	}
}
