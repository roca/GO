package main

import "slices"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)

	slices.Sort(nums)

	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2
	}

	return float64(nums[len(nums)/2])
}
