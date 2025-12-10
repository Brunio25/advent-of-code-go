package geom

type Coordinates3D struct {
	X int
	Y int
	Z int
}

func (c Coordinates3D) WithX(x int) Coordinates3D {
	return Coordinates3D{X: x, Y: c.Y, Z: c.Z}
}

func (c Coordinates3D) WithY(y int) Coordinates3D {
	return Coordinates3D{X: c.X, Y: y, Z: c.Z}
}

func (c Coordinates3D) WithZ(z int) Coordinates3D {
	return Coordinates3D{X: c.X, Y: c.Y, Z: z}
}

func (c Coordinates3D) EuclideanDistance(oc Coordinates3D) float64 {
	return euclideanDistanceHelper([]float64{
		float64(c.X - oc.X),
		float64(c.Y - oc.Y),
		float64(c.Z - oc.Z),
	})
}
