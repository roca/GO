package main

import "fmt"

func main() {
	convert("PAYPALISHIRING", 3)
}

func convert(s string, numRows int) string {
	m := matrix(s, numRows)

	for j := 0; j < len(s); j++ {
		if j%2 == 0 {
			for i := 0; i < numRows; i++ {
				m[i][j] = rune(s[j])
			}
		} else {
			for i := numRows - 1; i >= 0; i-- {
				m[i][j] = rune(s[j])
			}
		}
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < len(s); j++ {
			fmt.Print(string(m[i][j]))
		}
		fmt.Println()
	}
	return ""
}

func matrix(s string, numRows int) [][]rune {
	m := make([][]rune, numRows)
	for i := range m {
		m[i] = make([]rune, len(s))
	}
	return m
}
