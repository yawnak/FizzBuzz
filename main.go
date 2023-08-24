package main

import (
	"fmt"
	"strconv"
)

const (
	Fizz     = "Fizz"
	Buzz     = "Buzz"
	FizzBuzz = "FizzBuzz"
)

type FizzBuzzer interface {
	Fizz(n int) string
}

type FizzBuzzClassic struct {
}

func (FizzBuzzClassic) Fizz(n int) string {
	answerList := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		var answer string
		if i%15 == 0 {
			answer = FizzBuzz
		} else if i%3 == 0 {
			answer = Fizz
		} else if i%5 == 0 {
			answer = Buzz
		} else {
			answer = strconv.Itoa(i)
		}
		answerList = append(answerList, answer)
		answer = ""
	}
	return fmt.Sprint(answerList)
}
