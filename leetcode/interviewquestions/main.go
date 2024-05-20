package main


func removeDuplicates(nums []int) int {
	occurrences := occurrenceCount(nums)
	for i, v := range nums {
		if occurrences[v] > 2 {
			for occurrences[v] > 2 {
				nums =  append(nums[:i], nums[i+1:]...)
				occurrences[v]--
			}
		}
	}

	return len(nums)
}

func occurrenceCount(nums []int) map[int]int {
	occurrence := make(map[int]int)
	for _, num := range nums {
		occurrence[num]++
	}
	return occurrence
}
