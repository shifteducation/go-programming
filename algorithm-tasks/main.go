package main

import "fmt"

func main() {
	// nums = [2,7,11,15], target = 9
	//idxs := twoSum([]int{2, 7, 11, 15}, 9)
	//fmt.Println(idxs)
	//
	//// nums = [2,3,2,2,3], val = 3 - [2,2,_,_]
	//ints := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//result := removeElement(ints, 2)
	//fmt.Println(result)
	//fmt.Println(ints)
	//
	//fmt.Println(singleNumber([]int{4, 1, 2, 1, 2, 4, 5}))
	//
	//fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2, 2}))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	//fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
