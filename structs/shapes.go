package shapes

// Perimeter takes width and height of a rectangle and
// returns its perimeter.
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

// Area takes width and height of a rectangle and
// returns its area.
func Area(width, height float64) float64 {
	return width * height
}
