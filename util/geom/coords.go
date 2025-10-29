package geom

type Coordinates struct {
	X int
	Y int
}

func (c Coordinates) WithX(x int) Coordinates {
	return Coordinates{X: x, Y: c.Y}
}

func (c Coordinates) WithY(y int) Coordinates {
	return Coordinates{X: c.X, Y: y}
}
