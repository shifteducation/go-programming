package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"time"
)

type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string {
	return fmt.Sprintf("math: square root of negative number %g", float64(f))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, NegativeSqrtError(f)
	}
	return math.Sqrt(f), nil
}

type CustomError struct {
	message   string
	timestamp time.Time
}

func (c CustomError) Error() string {
	return fmt.Sprintf("%v: %s", c.timestamp, c.message)
}

const badInput = "abc"

type BadInputError struct {
	input string
}

func (e BadInputError) Error() string {
	return fmt.Sprintf("bad input: %s", e.input)
}

func main() {
	handleSqrt()
	checkIs()
	checkAs()
	fmt.Println(checkDefer(1))
}

func handleSqrt() {
	result, err := Sqrt(-1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func checkIs() {
	ErrInvalidValue := errors.New("invalid  value")
	err := fmt.Errorf("wrapped:  %w", ErrInvalidValue)
	fmt.Println("isErrInvalidValue: ", errors.Is(err, ErrInvalidValue))

	// -------
	sqrtError := NegativeSqrtError(-2)
	err = fmt.Errorf("wrapped:  %w", sqrtError)
	fmt.Println("isNegativeSqrtError: ", errors.Is(err, sqrtError))

	// -------
	customError := CustomError{"something is wrong", time.Now()}
	err = fmt.Errorf("wrapped:  %w", customError)
	fmt.Println("isNegativeSqrtError: ", errors.Is(err, customError))
}

func checkAs() {
	input := badInput

	err := fmt.Errorf("wrapped: %w", BadInputError{input: input})
	var badInputErr BadInputError
	if errors.As(err, &badInputErr) {
		fmt.Printf("bad input error occured: %s\n", badInputErr)
	}
}

func checkDefer(i int) int {
	defer func() int {
		i++
		return i
	}() // todo check
	return i
}
