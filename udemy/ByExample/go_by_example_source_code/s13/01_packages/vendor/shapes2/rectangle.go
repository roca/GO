package shapes2

import "fmt"

// Area ... your comments
func Area(l, w int) int {
	return l * w * 2 //this formula is incorrect; only for demonstration purposes;
}

func init() {
	fmt.Println("=> init() from shapes2.rectangle package.")
}
