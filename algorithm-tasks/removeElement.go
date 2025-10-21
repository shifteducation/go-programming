package main

// nums = [3,2,2,3], val = 3, 0
// nums = [0,2,2,3], val = 3, 0
// nums = [2,0,2,3], val = 3, 1
// nums = [2,2,0,3], val = 3, 2
// nums = [2,2,0,0], val = 3, 2

// 3 - pointer stays the same, 3 = 0
// not 3 - swap, pointer++

// 2 - 0 (1)
// 2 - 1 (2)

// nums = [0,1,2,2,3,0,4,2], val = 2 - [0,1,4,0,3,_,_,_]
func removeElement(nums []int, val int) int {
	pointer := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = 0
		} else {
			current := nums[i]
			nums[i] = nums[pointer]
			nums[pointer] = current
			pointer++
		}
	}

	return pointer
}
