package raindrops

const testVersion = 2

// Convert tranforms int to custom message
func Convert(n int) string {
	factors := [3]int{3, 5, 7}
	phrase := ""
	for _, factor := range factors {
		if n%factor == 0 && factor == 3 {
			phrase = phrase + "Pling"
		}
	}
	return phrase
}

// The test program has a benchmark too.  How fast does your Convert convert?
