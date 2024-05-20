package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 1, 2, 2, 3}
	expected := 5

	if got := removeDuplicates(nums); got != expected {
		t.Errorf("removeDuplicates(%v) = %v, want %v", nums, got, expected)
	}
}

func TestOccurrenceCount(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	expected := map[int]int{1: 3, 2: 2, 3: 1}

	// Map values are deeply equal when all of the following are true: 
	// they are both nil or both non-nil, they have the same length, 
	// and either they are the same map object or their corresponding 
	// keys (matched using Go equality) map to deeply equal values.

	if got := occurrenceCount(nums); !reflect.DeepEqual(got, expected) {
		t.Errorf("occurrenceCount(%v) = %v, want %v", nums, got, expected)
	}
}
