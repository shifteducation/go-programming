package main

import "fmt"

func main() {
	arrayExamples()
	sliceExamples()
}

func arrayExamples() {
	var arr1 [256]int // фиксированная длина
	fmt.Println(arr1)

	var arr2 [10][10]string // может быть многомерным
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3} //  автоматический подсчет длины
	fmt.Println(arr3)

	arr4 := [10]int{1, 2, 3, 4, 5}
	fmt.Println(arr4)

	ints := arr4[2:4] // получение среза
	fmt.Println(ints)
}

func sliceExamples() {
	ints1 := make([]int, 3)
	fmt.Println(ints1)

	ints2 := []int{1, 2, 3}
	ints3 := make([]int, 3)
	copy(ints3, ints2)
	fmt.Println(ints3)
	ints3[0] = 4
	fmt.Println(ints2, ints3)

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := s[3:5]
	fmt.Println(s2)
	s2 = s[3:]
	fmt.Println(s2)
	s2 = s[:5]
	fmt.Println(s2)
	s2 = s[:] // копия s (shallow)
	fmt.Println(s2)
	s2[0] = 1
	fmt.Println(s, s2)

	s = append(s[:2], s[2+1:]...)
	fmt.Println(s)

	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sl2 := make([]int, 2)
	copy(sl2, sl)
	sl2[0] = 1
	fmt.Println(sl, sl2)

	ints4 := make([]int, 3, 4)
	ints5 := append(ints4, 1)
	ints5[0] = 1
	fmt.Println(ints4, ints5) // 1 0 0, 1 0 0 1

	ints6 := append(ints5, 2)
	ints6[0] = 2
	fmt.Println(ints5, ints6)

	// Индекс и значение
	for i, v := range ints2 {
		fmt.Println(i, v)
	}

	// Только индекс
	for i := range ints2 {
		fmt.Println(i)
	}

	// Только значение
	for _, v := range ints2 {
		fmt.Println(v)
	}
}
