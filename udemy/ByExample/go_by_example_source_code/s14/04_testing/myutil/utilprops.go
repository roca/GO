package myutil

// Area ...
func Area(l, w uint) uint {
	return l * w
}

// Perim ...
func Perim(l, w uint) uint {
	return 2 * (l + w)
}

// Avg ...
func Avg(numbers []float32) float32 {
	total := float32(0)
	for _, n := range numbers {
		total += n
	}
	return total / float32(len(numbers))
}
