package main

import "testing"



func BenchmarkSum(b *testing.B) {
	// Create a slice of five integersA
	length := b.N
	nums := []int{}
	for i := 0; i < length; i++ {
		nums = append(nums, i)
	}
	sum := 0
	for n := 0; n  < len(nums) ; n++ { // This is not bound checked and runs faster
	// for n := 0; n < length; n++ {  // This is bounds checked and take longer to run
		// Sum over the array backwards
		sum += nums[n]
		_ = sum
	}
}
