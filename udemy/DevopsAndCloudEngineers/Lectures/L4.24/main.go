package main

import "fmt"

func main() {
	var ar1 [7]int = [7]int{7, 3, 6, 0, 4, 9, 10}
	fmt.Println("ar1",ar1)
	fmt.Printf("len: %d, cap: %d\n", len(ar1), cap(ar1))

	var ar2 []int = ar1[1:3]
	fmt.Println("ar2",ar2)
	fmt.Printf("len: %d, cap: %d\n", len(ar2), cap(ar2))

	ar2  = ar2[0: len(ar2) + 2]
	fmt.Println("ar2",ar2)
	fmt.Printf("len: %d, cap: %d\n", len(ar2), cap(ar2))

	for k := range ar2 {
		ar2[k]+=1
	}
	fmt.Println("ar2",ar2)
	fmt.Printf("len: %d, cap: %d\n", len(ar2), cap(ar2))
	fmt.Println("ar1",ar1)


	var ar3 []int  = []int{1,2,3}
	fmt.Println("ar3",ar3)
	fmt.Printf("len: %d, cap: %d\n", len(ar3), cap(ar3))
	ar3 = append(ar3, 4)
	fmt.Println("ar3",ar3)
	fmt.Printf("len: %d, cap: %d\n", len(ar3), cap(ar3))
	ar3 = append(ar3, 5)
	fmt.Println("ar3",ar3)
	fmt.Printf("len: %d, cap: %d\n", len(ar3), cap(ar3))
	ar3 = append(ar3, 6)
	fmt.Println("ar3",ar3)
	fmt.Printf("len: %d, cap: %d\n", len(ar3), cap(ar3))
	ar3 = append(ar3, 7)
	fmt.Println("ar3",ar3)
	fmt.Printf("len: %d, cap: %d\n", len(ar3), cap(ar3))


	ar4 := make([]int,3,9)
	fmt.Println("ar4",ar4)
	fmt.Printf("len: %d, cap: %d\n", len(ar4), cap(ar4))
}
