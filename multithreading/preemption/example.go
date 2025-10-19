package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		var u int16
		for {
			u -= 2
			if u == 1 { // never happens
				break
			}
		}
	}()
	t := <-time.After(time.Millisecond * 5)
	fmt.Println(t)

	fmt.Println("go 1.13 has never been here")
}
