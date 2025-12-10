package geom_test

import (
	"advent-of-code-go/util/geom"
	"fmt"
	"testing"
)

func TestCoordinates3D_WithX(t *testing.T) {
	original := geom.Coordinates3D{X: 0, Y: 0, Z: 0}
	expected := geom.Coordinates3D{X: 27, Y: 0, Z: 0}
	if actual := original.WithX(27); actual != expected {
		t.Errorf("coordinates.Coordinates3d.WithX() = Coordinates3D%v, Expected Coordinates3D%v", actual, expected)
	}
}

func TestCoordinates3D_WithY(t *testing.T) {
	original := geom.Coordinates3D{X: 0, Y: 0, Z: 0}
	expected := geom.Coordinates3D{X: 0, Y: 27, Z: 0}
	if actual := original.WithY(27); actual != expected {
		t.Errorf("coordinates.Coordinates3d.WithY() = Coordinates3D%v, Expected Coordinates3D%v", actual, expected)
	}
}

func TestCoordinates3D_WithZ(t *testing.T) {
	original := geom.Coordinates3D{X: 0, Y: 0, Z: 0}
	expected := geom.Coordinates3D{X: 0, Y: 0, Z: 27}
	if actual := original.WithZ(27); actual != expected {
		t.Errorf("coordinates.Coordinates3d.WithZ() = Coordinates3D%v, Expected Coordinates3D%v", actual, expected)
	}
}

func TestCoordinates3D_EuclideanDistance(t *testing.T) {
	testCases := []struct {
		p1       geom.Coordinates3D
		p2       geom.Coordinates3D
		expected float64
	}{
		{
			p1:       geom.Coordinates3D{X: 1, Y: 1, Z: 1},
			p2:       geom.Coordinates3D{X: 10, Y: 10, Z: 10},
			expected: 15.588457268119896,
		},
		{
			p1:       geom.Coordinates3D{X: 3, Y: -1, Z: 30},
			p2:       geom.Coordinates3D{X: 12, Y: 57, Z: -53},
			expected: 101.65628362280415,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`"%v-%v"->%v`, tc.p1, tc.p2, tc.expected), func(t *testing.T) {
			if actual := tc.p1.EuclideanDistance(tc.p2); actual != tc.expected {
				t.Errorf("coordinates.Coordinates3d().EuclidenaDistance() = %v, Expected %v", actual, tc.expected)
			}
		})
	}
}
