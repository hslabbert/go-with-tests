package shapes

import "math"

// A Rectangle has Height and Width to represent a rectangle shape.
type Rectangle struct {
	Width  float64
	Height float64
}

// A Circle has Radius to represent a circle shape.
type Circle struct {
	Radius float64
}

// Perimeter takes width and height of a rectangle and
// returns its perimeter.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns the area of a Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area returns the area of a Circle
func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}
