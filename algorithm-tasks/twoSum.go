package main

import "errors"

func twoSum(nums []int, target int) ([]int, error) {
	if len(nums) < 2 {
		return nil, errors.New("len of nums must be greater than 2")
	}

	valueToIdx := map[int]int{}
	result := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		current := nums[i]

		if v, ok := valueToIdx[target-current]; ok == true {
			result[0] = v
			result[1] = i
			return result, nil
		} else {
			valueToIdx[current] = i
		}
	}

	return nil, errors.New("not found")
}
