package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 4 {
			break
		}

		sum += i
	}

	fmt.Println(sum)

	var i int
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
