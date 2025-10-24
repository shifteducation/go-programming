package main

type IntSet map[int]struct{}

func findIntersection(nums1, nums2 []int) []int {
	var set = make(IntSet)
	result := make([]int, 0)
	var numsForSet []int
	var nums []int
	if len(nums1) >= len(nums2) {
		numsForSet = nums1
		nums = nums2
	} else {
		numsForSet = nums2
		nums = nums1
	}

	for _, num := range numsForSet {
		set[num] = struct{}{}
	}

	for _, num := range nums {
		if _, ok := set[num]; ok {
			result = append(result, num)
		}
	}

	return result
}
