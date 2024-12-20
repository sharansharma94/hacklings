package arrays

import (
	"reflect"
	"testing"
)

// Test cases for FindMaxNumber
func TestFindMaxNumber(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3, 2, 5, 4}, 5},
		{[]int{1}, 1},
		{[]int{}, -1},
		{[]int{-1, -5, -2, -8}, -1},
		{[]int{10, 10, 10}, 10},
	}

	for _, test := range tests {
		result := FindMaxNumber(test.nums)
		if result != test.expected {
			t.Errorf("For nums=%v, expected %d but got %d",
				test.nums, test.expected, result)
		}
	}
}

// Helper function to compare two slices
func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum(%v, %v) = %v; want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
