package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

// Simple test of FizzBuzzModular
func TestFizzBuzzModularClassic(t *testing.T) {
	//case struct
	type Case struct {
		input  int
		output string
	}
	//test data
	cases := []Case{
		{1, "1"},
		{3, Fizz},
		{5, Buzz},
		{15, FizzBuzz},
	}

	fb := NewFizzBuzzModular(fzbzClassic, fzClassic, bzClassic) //create classic fizzbuzzer

	//check fizzbuzzer against test data
	for _, c := range cases {
		assert := fb.FizzBuzz(c.input)
		if assert != c.output {
			t.Errorf("wrong output fizzbuzz expected %s, got %s", c.output, assert)
		}
	}
}

// Simple test of incremental FizzBuzzIterator
func TestIncrementalFizzBuzzIterator(t *testing.T) {
	//test data
	input := 20
	output := []string{
		"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
		"16", "17", "Fizz", "19", "Buzz",
	}

	//create incrementalfizzbuzziterator
	fb := NewFizzBuzzModular(fzbzClassic, fzClassic, bzClassic)
	fbIterator := NewIncrementalFizzBuzzIterator(fb)

	//check output slice against test data
	fbIterator.Iterate(input)
	for i, v := range output {
		//check if iterator has next value
		if !fbIterator.Next() {
			t.Errorf("error, iterator has no value when expected %s", v)
		} else if fbIterator.Scan() != v {
			t.Errorf("error on output #%d, expected %s, got %s", i, v, fbIterator.Scan())
		}
	}
}

// Test that uses 100 test points
func TestFizzBuzzModularClassicWithData(t *testing.T) {
	// Open the file
	file, err := os.Open("test_data/test1.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		_, err = strconv.Atoi(firstLine)
		if err != nil {
			fmt.Println("error converting to integer:", err)
			return
		}
	} else {
		fmt.Println("error reading first line:", scanner.Err())
		return
	}

	// Iterate over each line and compare it with FizzBuzzer output
	i := 1
	fb := NewFizzBuzzModular(fzbzClassic, fzClassic, bzClassic) //create classic fizzbuzzer
	for scanner.Scan() {
		output := scanner.Text()
		assert := fb.FizzBuzz(i)
		//fmt.Println(output, assert)
		if output != assert {
			t.Errorf("wrong value, expected %s, got %s", output, assert)
		}
		i++
	}

	if scanner.Err() != nil {
		fmt.Println("error reading lines:", scanner.Err())
	}
}

func TestIncrementalFizzBuzzIteratorWithData(t *testing.T) {
	// Open the file
	file, err := os.Open("test_data/test1.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	var n int
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		n, err = strconv.Atoi(firstLine)
		if err != nil {
			fmt.Println("error converting to integer:", err)
			return
		}
	} else {
		fmt.Println("error reading first line:", scanner.Err())
		return
	}
	//create iterator
	fb := NewFizzBuzzModular(fzbzClassic, fzClassic, bzClassic) //classic fizzbuzzer
	fbIterator := NewIncrementalFizzBuzzIterator(fb)
	fbIterator.Iterate(n) //do iterations

	// Iterate over each line and print
	i := 0
	for scanner.Scan() {
		output := scanner.Text()
		if !fbIterator.Next() {
			t.Errorf("iterator has no value, expected %s", output)
		} else if output != fbIterator.Scan() {
			t.Errorf("wrong value, expected %s, got %s", output, fbIterator.Scan())
		}
		i++
	}

	if scanner.Err() != nil {
		fmt.Println("error reading lines:", scanner.Err())
	}
}
