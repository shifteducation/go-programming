package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	fmt.Println("settability of p:", p.CanSet())
	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(x)

	//var x uint8 = 'x'
	//fmt.Println("value:", reflect.ValueOf(x))
	//fmt.Println("type:", reflect.TypeOf(x))
	//v := reflect.ValueOf(&x)
	//fmt.Println("type:", v.Type())                            // uint8.
	//fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	//x = uint8(v.Elem().Uint())                                // v.Uint returns a uint64.
	//fmt.Println(v.Elem().CanSet())
	//v.Elem().SetUint(1)
	//fmt.Println(x)
}
