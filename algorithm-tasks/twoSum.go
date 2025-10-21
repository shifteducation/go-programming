package main

func twoSum(nums []int, target int) []int {
	valueToIdx := map[int]int{}
	result := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		current := nums[i]

		if v, ok := valueToIdx[target-current]; ok == true {
			result[0] = v
			result[1] = i
			return result
		} else {
			valueToIdx[current] = i
		}
	}

	return result
}
