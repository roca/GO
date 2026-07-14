package main

import (
	"crypto/sha256"
	"fmt"
)

func checksumMatches(message string, checksum string) bool {
	h := sha256.New()
	h.Write([]byte(message))
	return fmt.Sprintf("%x", h.Sum(nil)) == checksum
}

// don't touch below this line

func test(message, checksum string) {
	fmt.Printf("Checking message '%v'...\n", message)
	fmt.Printf("Expected checksum: %v\n", checksum)
	if checksumMatches(message, checksum) {
		fmt.Println("Checksum matches!")
	} else {
		fmt.Println("Checksum does not match!")
	}
	fmt.Println("========")
}

func main() {
	test("pa$$w0rd", "4b358ed84b7940619235a22328c584c7bc4508d4524e75231d6f450521d16a17")
	test("buil4WithB1ologee", "1c489a153271aaf3b234aa154b1a2eef5248eb9ab402e4d3c8b7bc3d81fed1a8")
	test("br3ak1ngB@d1sB3st", "5d178e1c6fd5d76415e1632f84e5192fb50ef244d42a02148fedbf991d914546")
	test("b3ttterC@llS@ulI$B3tter", "8d42f2dc81476123974619969a42b27b8d8a4fa507be99c9623f614ad2d859f7")
}
