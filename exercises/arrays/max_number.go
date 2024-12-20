package arrays

// FindMaxNumber returns the largest number in the given array
// If the array is empty, returns -1
func FindMaxNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}
