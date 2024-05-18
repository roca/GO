package main

import "fmt"

func main() {

	nums := []int{1, 1, 1, 2, 2, 3}

	fmt.Println(nums)

	fmt.Println(removeDuplicates(nums))

}

func removeDuplicates(nums []int) int {
	occurrences := occurrenceCount(nums)
	fmt.Println(occurrences)
	for i, v := range nums {
		if occurrences[v] > 2 {
			nums = append(nums[:i], nums[i+1:]...)
			occurrences[v]--
		}
	}

	fmt.Println("After", nums)
	return len(nums)
}

func occurrenceCount(nums []int) map[int]int {
	occurrence := make(map[int]int)
	for _, num := range nums {
		occurrence[num]++
	}
	return occurrence
}
