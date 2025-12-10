package geom

type Coordinates2D struct {
	X int
	Y int
}

func (c Coordinates2D) WithX(x int) Coordinates2D {
	return Coordinates2D{X: x, Y: c.Y}
}

func (c Coordinates2D) WithY(y int) Coordinates2D {
	return Coordinates2D{X: c.X, Y: y}
}

func (c Coordinates2D) EuclideanDistance(oc Coordinates2D) float64 {
	return euclideanDistanceHelper([]float64{
		float64(c.X - oc.X),
		float64(c.Y - oc.Y),
	})
}
