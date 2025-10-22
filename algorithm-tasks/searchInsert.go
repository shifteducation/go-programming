package main

// 1, 3, 5, 6 - 0, 3, 1
// 1, 3, 5, 6 - 1, 3, 2
// 1, 3, 5, 6 - 2, 3, 2
func searchInsert(nums []int, target int) int {
	leftIndex := 0
	rightIndex := len(nums) - 1
	for leftIndex <= rightIndex {
		index := leftIndex + (rightIndex-leftIndex)/2
		// todo
		if nums[index] == target {
			return index
		}
		if target > nums[index] {
			leftIndex = index + 1
		} else {
			rightIndex = index - 1
		}
	}
	return len(nums)
}
