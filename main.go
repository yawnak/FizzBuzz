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

func (FizzBuzzClassic) Fizz(i int) string {
	if i%15 == 0 {
		return FizzBuzz
	} else if i%3 == 0 {
		return Fizz
	} else if i%5 == 0 {
		return Buzz
	}
	return strconv.Itoa(i)
}

type fizzbuzz struct {
	fzbz func(int) bool
	fz   func(int) bool
	bz   func(int) bool
}

func (fb *fizzbuzz) Fizz(i int) string {
	if fb.fzbz(i) {
		return FizzBuzz
	} else if fb.fz(i) {
		return Fizz
	} else if fb.bz(i) {
		return Buzz
	}
	return strconv.Itoa(i)
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
