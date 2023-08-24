package main

import "testing"

func TestSimpleTestFizzBuzz(t *testing.T) {
	type Case struct {
		input  int
		output string
	}
	cases := []Case{
		{1, "1"},
		{3, Fizz},
		{5, Buzz},
		{15, FizzBuzz},
	}

	fb := NewFizzBuzzModular(fzbzClassic, fzClassic, bzClassic)

	for _, c := range cases {
		assert := fb.FizzBuzz(c.input)
		if assert != c.output {
			t.Errorf("wrong output fizzbuzz expected %s, got %s", c.output, assert)
		}
	}
}
