package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func(ch chan<- int) {
		defer close(ch)
		time.Sleep(1 * time.Millisecond)
		ch <- 1
	}(ch1)

	for {
		select {
		case x := <-ch1:
			fmt.Println(x)
			//if !ok {
			//	ch1 = nil
			//}
		case t := <-time.NewTimer(2 * time.Second).C:
			fmt.Println(t)
		}
	}
}
