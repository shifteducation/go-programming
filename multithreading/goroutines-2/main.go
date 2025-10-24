package main

import (
	"fmt"
	fanoutfanin "shift/go-basics/goroutines/fan-out-fan-in"
	"sync"
)

func main() {
	checkClose()
	fanOutFanIn()
	checkMutexWithMap()
}

func checkMutexWithMap() {
	mutex := sync.Mutex{}
	ints := make(map[int]struct{})

	wg := sync.WaitGroup{}
	count := 10
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			defer mutex.Unlock()
			mutex.Lock()
			ints[i] = struct{}{}
		}(i)
	}
	wg.Wait()
	fmt.Println(len(ints))
}

func checkClose() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// close(ch) // try to comment - deadlock
	}()

	for i := range ch {
		fmt.Println(i)
	}
}

func fanOutFanIn() {
	nums := make([]int, 100)
	for i := range 100 {
		nums[i] = i
	}
	in := fanoutfanin.Generate(nums...)

	// Distribute the Square work across two goroutines that both read from in.
	c1 := fanoutfanin.Square(in)
	c2 := fanoutfanin.Square(in)

	// Consume the merged output from c1 and c2.
	for n := range fanoutfanin.Merge(c1, c2) {
		fmt.Println(n)
	}
}
