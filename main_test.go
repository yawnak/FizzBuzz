package main

import "testing"

func TestFizzBuzz(t *testing.T) {
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

func TestIncrementalFizzBuzzIterator(t *testing.T) {
	input := 20
	output := []string{
		"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
		"16", "17", "Fizz", "19", "Buzz",
	}

	fb := NewFizzBuzzModular(fzbzClassic, fzClassic, bzClassic)
	fbIterator := NewIncrementalFizzBuzzIterator(fb)

	assert := fbIterator.Iterate(input)
	for i, v := range output {
		if assert[i] != v {
			t.Errorf("error on output #%d, expected %s, got %s", i, v, assert[i])
		}
	}
}
