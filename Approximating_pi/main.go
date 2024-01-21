package main

import (
	"fmt"
	"pi/pi"
)

func main() {
	// n is the number of sets of coordinates we generate to estimate pi
	n := 10000
	// hits is the amount of coordinates that fell inside a unit circle
	hits := 0

	for i := 0; i < n; i++ {
		x, y := pi.GenerateXY()
		if pi.IsCoordinateInsideUnitCircle(x, y) {
			hits++
		}
	}

	/*
		The maths:
		let:
		r = radius of unit circle
		l = length of unit square
		(so l = 2r)

		A1) Area of circle=pi*r^2
		A2) Area of square=l^2

		The ratio A1/A2=pi/4 (with substitutions above)
		=>multiply our ratio by 4 for an approximation of pi.
	*/
	approx := (float64(hits) / float64(n)) * 4

	fmt.Printf("Pi is approximately %f (n=%d)\n", approx, n)
}
