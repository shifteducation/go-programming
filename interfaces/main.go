package main

import (
	"fmt"
	"strconv"
)

type Temperature struct {
	Temp int
}

func (t Temperature) String() string {
	return strconv.Itoa(t.Temp) + " C"
}

func main() {
	x := Temperature{24}
	fmt.Printf("%v", x)
}
