package main

const (
	Fizz     = "Fizz"
	Buzz     = "Buzz"
	FizzBuzz = "FizzBuzz"
)

type FizzBuzzer interface {
	Fizz(n int) string
}


