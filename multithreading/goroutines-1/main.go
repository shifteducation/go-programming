package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var once = sync.Once{}

func main() {
	checkSelect()
	checkWaitGroup()
	checkMutex()
	checkSyncMap()
	checkAtomic()
	checkSyncOnce()
}

func checkWaitGroup() {
	wg := sync.WaitGroup{}
	count := 5
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("go")
		}()
	}
	wg.Wait()
}

func checkSelect() {
	ch1 := make(chan int)

	go func(ch chan<- int) {
		defer close(ch)
		time.Sleep(1 * time.Millisecond)
		ch <- 1
	}(ch1)

	for {
		select {
		case x, ok := <-ch1:
			fmt.Println(x)
			if !ok {
				ch1 = nil
			}
		case t := <-time.NewTimer(2 * time.Second).C:
			fmt.Println(t)
		}
	}
}

func checkMutex() {
	wg := sync.WaitGroup{}
	count := 1000
	wg.Add(count)

	v := 0
	var mutex sync.Mutex

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			defer mutex.Unlock()

			mutex.Lock()
			v++
		}()
	}
	wg.Wait()
	fmt.Println(v)
}

func checkSyncMap() {
	m := sync.Map{}
	wg := sync.WaitGroup{}
	count := 100
	wg.Add(count)

	for i := range count {
		go func(x int) {
			defer wg.Done()
			m.Store(x, x)
		}(i)
	}

	wg.Wait()

	m.Range(func(k, v any) bool {
		fmt.Println(k, v)
		return true
	})
}

func checkAtomic() {
	var x int32
	wg := sync.WaitGroup{}
	count := 100
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&x, 1)
		}()
	}

	wg.Wait()
	fmt.Println(x)
}

func checkSyncOnce() {
	doSomething := func() int {
		result := 0
		once.Do(
			func() {
				fmt.Println("sync.Once")
				result++
			},
		)
		return result
	}

	wg := sync.WaitGroup{}
	count := 100
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(doSomething())
		}()
	}

	wg.Wait()
}
