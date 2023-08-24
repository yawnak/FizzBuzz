package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// Fizz Buzz constants
const (
	Fizz     = "Fizz"
	Buzz     = "Buzz"
	FizzBuzz = "FizzBuzz"
)

type FizzBuzzer interface {
	FizzBuzz(n int) string
}

// FizzBuzzModular recieves 3 functions that will determine output of FizzBuzz
type FizzBuzzModular struct {
	//fzbz recieves int and returns bool if int is FizzBuzz
	fzbz func(int) bool
	//fz recieves int and returns bool if int is Fizz
	fz func(int) bool
	//bz recieves int and returns bool if int is Buzz
	bz func(int) bool
}

// FizzBuzz implements FizzBuzzer interface
// It checks i against fzbz, fz, and bz and returns FizzBuzz, Fizz or Buzz respectively.
// Or it returns i as a string if all checks fail.
func (fb FizzBuzzModular) FizzBuzz(i int) string {
	if fb.fzbz(i) {
		return FizzBuzz
	} else if fb.fz(i) {
		return Fizz
	} else if fb.bz(i) {
		return Buzz
	}
	return strconv.Itoa(i)
}

// FizzBuzzModular constructor. If a function is nil, it will replace it with a function that always returns false.
func NewFizzBuzzModular(fzbz func(int) bool, fz func(int) bool, bz func(int) bool) FizzBuzzModular {
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
	return FizzBuzzModular{
		fzbz: fzbz,
		fz:   fz,
		bz:   bz,
	}
}

// classic implementation of fzbz func. Returns true if i is divisible by 15
func fzbzClassic(i int) bool {
	return i%15 == 0
}

// classic implementation of fz function. Returns true if i is divisible by 3
func fzClassic(i int) bool {
	return i%3 == 0
}

// classic implementation of bz function. Returns true if i is divisible by 5
func bzClassic(i int) bool {
	return i%5 == 0
}

type Iteratorer interface {
	Iterate(int)
	Scan() string
	Next() bool
}

type IncrementalFizzBuzzIterator struct {
	curidx int
	n      int
	curval string
	fb     FizzBuzzer
}

func (iterator *IncrementalFizzBuzzIterator) Iterate(n int) {
	iterator.curidx = 1
	iterator.n = n
}

func (iterator *IncrementalFizzBuzzIterator) Next() bool {
	if iterator.curidx > iterator.n {
		return false
	}
	iterator.curval = iterator.fb.FizzBuzz(iterator.curidx)
	iterator.curidx += 1
	return true
}

func (iterator IncrementalFizzBuzzIterator) Scan() string {
	return iterator.curval
}

func NewIncrementalFizzBuzzIterator(fb FizzBuzzer) IncrementalFizzBuzzIterator {
	return IncrementalFizzBuzzIterator{
		fb: fb,
	}
}

var (
	n int
)

func main() {
	flag.IntVar(&n, "n", 0, "input n, default 0")
	flag.Parse()

	fb := NewFizzBuzzModular(fzbzClassic, fzClassic, bzClassic)
	fbIncremental := NewIncrementalFizzBuzzIterator(fb)
	fbIncremental.Iterate(n)

	builder := strings.Builder{}
	for fbIncremental.Next() {
		builder.WriteString(fbIncremental.Scan())
		builder.WriteString(",")
	}
	fmt.Print(builder.String()[:builder.Len()-1])
}
