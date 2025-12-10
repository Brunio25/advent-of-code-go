package grid

import "advent-of-code-go/util/geom"

type Direction func(c geom.Coordinates2D) geom.Coordinates2D

func Up(c geom.Coordinates2D) geom.Coordinates2D {
	return geom.Coordinates2D{X: c.X, Y: c.Y - 1}
}

func Down(c geom.Coordinates2D) geom.Coordinates2D {
	return geom.Coordinates2D{X: c.X, Y: c.Y + 1}
}

func Left(c geom.Coordinates2D) geom.Coordinates2D {
	return geom.Coordinates2D{X: c.X - 1, Y: c.Y}
}

func Right(c geom.Coordinates2D) geom.Coordinates2D {
	return geom.Coordinates2D{X: c.X + 1, Y: c.Y}
}

func LeftUp(c geom.Coordinates2D) geom.Coordinates2D {
	return geom.Coordinates2D{X: c.X - 1, Y: c.Y - 1}
}

func RightUp(c geom.Coordinates2D) geom.Coordinates2D {
	return geom.Coordinates2D{X: c.X + 1, Y: c.Y - 1}
}
func LeftDown(c geom.Coordinates2D) geom.Coordinates2D {
	return geom.Coordinates2D{X: c.X - 1, Y: c.Y + 1}
}

func RightDown(c geom.Coordinates2D) geom.Coordinates2D {
	return geom.Coordinates2D{X: c.X + 1, Y: c.Y + 1}
}

var AllDirections = []Direction{Up, RightUp, Right, RightDown, Down, LeftDown, Left, LeftUp}
