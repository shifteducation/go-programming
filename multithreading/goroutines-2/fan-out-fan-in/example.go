package fan_out_fan_in

import "sync"

func Generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}

// The merge function converts a list of channels to a single channel by starting a goroutine for each inbound channel
// that copies the values to the sole outbound channel.
// Once all the output goroutines have been started, merge starts one more goroutine to close the outbound channel after all sends on that channel are done.
func Merge(chs ...<-chan int) <-chan int {
	out := make(chan int)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(chs))

	for _, ch := range chs {
		go func() {
			defer waitGroup.Done()
			for num := range ch {
				out <- num
			}
		}()
	}

	go func() {
		waitGroup.Wait()
		close(out)
	}()

	return out
}
