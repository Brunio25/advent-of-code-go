package geom

import "math"

func euclideanDistanceHelper(coordDiffs []float64) float64 {
	squaresSum := float64(0)
	for _, diff := range coordDiffs {
		squaresSum += math.Pow(diff, 2)
	}

	return math.Sqrt(squaresSum)
}
