// File name: ...\s10\08_interface_sort_struct\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"sort"
)

type person []string

func (p person) Len() int           { return len(p) }
func (p person) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p person) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	s := person{"Susan", "Tyler", "Ava", "Nick"}
	fmt.Println(s)

	sort.Sort(s)
	fmt.Println(s)

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
