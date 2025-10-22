package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x uint8 = 'x'
	v := reflect.ValueOf(&x)
	fmt.Println("type:", v.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	x = uint8(v.Elem().Uint())                                // v.Uint returns a uint64.
	fmt.Println(v.Elem().CanSet())
	v.Elem().SetUint(1)
	fmt.Println(x)
}
