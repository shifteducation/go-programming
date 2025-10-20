package main

import "fmt"

func main() {
	//При присвоении переменных типа структура - данные копируются.

	//a := struct{ x, y int }{0, 0}
	//b := a
	//a.x = 1
	//
	//fmt.Println(b.x) // ?

	//При присвоении указателей - копируется только адрес данных.

	//a := &struct{ x, y int }{}
	//b := a
	//a.x = 1
	//
	//fmt.Println(b.x) // ?

	a := struct{ x *int }{new(int)}
	b := a
	*a.x = 1

	fmt.Println(*b.x) // ?
}
