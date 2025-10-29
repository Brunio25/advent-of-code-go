package geom_test

import (
	"advent-of-code-go/util/geom"
	"testing"
)

func TestWithX(t *testing.T) {
	original := geom.Coordinates{X: 0, Y: 0}
	expected := geom.Coordinates{X: 27, Y: 0}
	if actual := original.WithX(27); actual.X != 27 || actual.Y != 0 {
		t.Errorf("coordinates.WithX() = Coordinates%v, Expected Coordinates%v", actual, expected)
	}
}

func TestWithY(t *testing.T) {
	original := geom.Coordinates{X: 0, Y: 0}
	expected := geom.Coordinates{X: 0, Y: 27}
	if actual := original.WithY(27); actual.X != 0 || actual.Y != 27 {
		t.Errorf("coordinates.WithX() = Coordinates%v, Expected Coordinates%v", actual, expected)
	}
}
