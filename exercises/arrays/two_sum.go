package arrays

// TwoSum finds two numbers in the array that add up to the target
func TwoSum(nums []int, target int) []int {
	// This is where you'll implement your solution

	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, found := seen[complement]; found {
			return []int{index, i}
		}
		seen[num] = i
	}
	return nil
}
