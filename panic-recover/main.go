package main

import (
	"errors"
	"log"
)

func main() {
	doSomething()
}

func doSomething() {
	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
		}
	}()
	callPanic()
}

func callPanic() {
	err := errors.New("panic")
	panic(err)
}
