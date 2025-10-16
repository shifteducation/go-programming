package main

import (
	"errors"
	"fmt"
)

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	r, ok := i.(fmt.Stringer)
	fmt.Println(r, ok)

	var x error = errors.New("test error")
	t, ok := x.(error)
	fmt.Println(t, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // panic
	//fmt.Println(f)

	printType(i)
}

func printType(x any) {
	switch x.(type) {
	case string:
		fmt.Printf("%v has type string", x)
	case fmt.Stringer:
		fmt.Printf("%v has type fmt.Stringer", x)
	case error:
		fmt.Printf("%v has type error", x)
	case float64:
		fmt.Printf("%v has type float64", x)
	default:
		fmt.Printf("%v type is unknown", x)
	}
}
