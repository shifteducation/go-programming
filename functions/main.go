package main

import "fmt"

func countWithStep(step int) func() int {
	counter := 0

	return func() int { // эта функция является замыканием
		counter += step // и замыкает в себе переменную counter
		return counter
	}
}

func sumInts(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	increment := countWithStep(1)
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	//fmt.Println(sumInts(1, 2, 3, 4))
}
