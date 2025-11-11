package grid

import "advent-of-code-go/util/geom"

type Direction func(c geom.Coordinates) geom.Coordinates

func Up(c geom.Coordinates) geom.Coordinates {
	return geom.Coordinates{X: c.X, Y: c.Y - 1}
}

func Down(c geom.Coordinates) geom.Coordinates {
	return geom.Coordinates{X: c.X, Y: c.Y + 1}
}

func Left(c geom.Coordinates) geom.Coordinates {
	return geom.Coordinates{X: c.X - 1, Y: c.Y}
}

func Right(c geom.Coordinates) geom.Coordinates {
	return geom.Coordinates{X: c.X + 1, Y: c.Y}
}

func LeftUp(c geom.Coordinates) geom.Coordinates {
	return geom.Coordinates{X: c.X - 1, Y: c.Y - 1}
}

func RightUp(c geom.Coordinates) geom.Coordinates {
	return geom.Coordinates{X: c.X + 1, Y: c.Y - 1}
}
func LeftDown(c geom.Coordinates) geom.Coordinates {
	return geom.Coordinates{X: c.X - 1, Y: c.Y + 1}
}

func RightDown(c geom.Coordinates) geom.Coordinates {
	return geom.Coordinates{X: c.X + 1, Y: c.Y + 1}
}
