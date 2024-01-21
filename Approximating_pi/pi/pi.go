package pi

import (
	"math"
	"math/rand"
)

func GenerateXY() (float64, float64) {
	return rand.Float64(), rand.Float64()
}

func IsCoordinateInsideUnitCircle(x, y float64) bool {
	return math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)) <= 1
}
