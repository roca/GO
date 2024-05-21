package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {
	sbs := make([]string, numRows)

	j := 0
	for j < len(s) {

		for i := 0; i < numRows && j < len(s); i++ {
			sbs[i] += string(s[j])
			j++
		}

		for i := numRows - 2; i > 0 && j < len(s); i-- {
			sbs[i] += string(s[j])
			j++
		}

	}

	result := ""
	for i := 0; i < numRows; i++ {
		result += sbs[i]
	}
	return result
}
