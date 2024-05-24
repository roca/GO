package main

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	got := findMedianSortedArrays(nums1, nums2)
	want := 2.0
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	got = findMedianSortedArrays(nums1, nums2)
	want = 2.5
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
