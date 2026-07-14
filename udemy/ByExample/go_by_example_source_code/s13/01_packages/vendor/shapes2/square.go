package shapes2

import "fmt"

// SquareArea ... your comments
func SquareArea(s int) int {
	return s * s * 2 //this formula is incorrect; only for demonstration purposes;
}

func init() {
	fmt.Println("=> init() from shapes2.square package.")
}
