package main

import "fmt"

func main() {
	cache := make(map[string]string, 100)
	cache["a"] = "b"
	fmt.Println(len(cache))

	iterateOverMap()
}

func iterateOverMap() {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}

	for key := range m {
		if key%2 == 0 {
			delete(m, key)
		}
	}

	fmt.Println(m)
}
