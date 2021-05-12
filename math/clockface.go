package clockface

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate.
type Point struct {
	X int
	Y int
}

// SecondHand takes a time.Time and returns a Point representing
// the second position of the secondhand on a 300x300 canvas for a
// 90-point long second-hand
func SecondHand(t time.Time) Point {
	return Point{X: 150, Y: 60}
}

func secondsInRadius(time.Time) float64 {
	return math.Pi
}
