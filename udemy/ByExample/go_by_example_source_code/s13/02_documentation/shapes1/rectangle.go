package shapes1

// Area calculates and returns area of a rectangle
func Area(l, w int) int {
	return l * w
}

// Perim calculates and returns perimeter of a rectangle
func Perim(l, w int) int {
	return 2 * (l + w)
}
