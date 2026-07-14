// File name: ...\s10\07_interface_sort2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"sort"
)

type iSortable []int

func (p iSortable) Len() int           { return len(p) }
func (p iSortable) Less(i, j int) bool { return p[i] < p[j] }
func (p iSortable) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	n := []int{7, 2, 10, 14, 1, 14, 9}
	fmt.Println(n)

	sort.Sort(iSortable(n))
	// sort.Sort(sort.IntSlice(n))
	fmt.Println(n)

	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	s := []string{"Susan", "Tyler", "Ava", "Nick"}
	fmt.Printf("\n%s \n", s)

	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
