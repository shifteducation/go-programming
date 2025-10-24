package main

import "fmt"

//go:generate stringer -type=Color
type Color int

const (
	Red Color = iota
	Green
	Blue
)

func main() {
	var x Color = 9
	fmt.Println(x)
}
