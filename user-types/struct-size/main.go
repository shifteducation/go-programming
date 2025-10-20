package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := struct {
		a bool   // 1 (offset 0)
		b bool   // 1 (offset 1)
		c string // 16 (offset 8)
	}{
		true, true, "qqq",
	}
	fmt.Println(unsafe.Sizeof(x)) // 24!

	y := struct {
		c bool   // 1 (offset 0)
		b string // 16 (offset 1)
		a bool   // 1 (offset 24)
	}{
		true, "qqq", true,
	}
	fmt.Println(unsafe.Sizeof(y)) // 32!
}
