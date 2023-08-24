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

type fizzbuzz struct {
	fzbz func(int) bool
	fz   func(int) bool
	bz   func(int) bool
}

func (fb *fizzbuzz) Fizz(n int) string {
	answerList := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if fb.fzbz(i) {
			answerList = append(answerList, FizzBuzz)
		} else if fb.fz(i) {
			answerList = append(answerList, Fizz)
		} else if fb.bz(i) {
			answerList = append(answerList, Buzz)
		} else {
			answerList = append(answerList, strconv.Itoa(i))
		}
	}
	return fmt.Sprint(answerList)
}

func NewFizzBuzz(fzbz func(int) bool, fz func(int) bool, bz func(int) bool) fizzbuzz {
	if fzbz == nil {
		fzbz = func(i int) bool {
			return false
		}
	}
	if fz == nil {
		fz = func(i int) bool {
			return false
		}
	}
	if bz == nil {
		bz = func(int) bool {
			return false
		}
	}
	return fizzbuzz{
		fzbz: fzbz,
		fz:   fz,
		bz:   bz,
	}
}

func fzbzClassic(i int) bool {
	return i%15 == 0
}

func fzClassic(i int) bool {
	return i%3 == 0
}

func bzClassic(i int) bool {
	return i%5 == 0
}

func main() {
	n := 15
	fb := FizzBuzzClassic{}
	resClassic := fb.Fizz(n)

	fb2 := NewFizzBuzz(fzbzClassic, fzClassic, bzClassic)
	res2 := fb2.Fizz(n)

	fmt.Println(resClassic)
	fmt.Println(res2)
}
