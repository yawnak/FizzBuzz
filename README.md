Task description:  
Given an integer n, return and print a string array answer for (1 .. n)  
 - answer[i] == "FizzBuzz" if i is divisible by 3 and 5.  
 - answer[i] == "Fizz" if i is divisible by 3.  
 - answer[i] == "Buzz" if i is divisible by 5.  
 - answer[i] == i (as a string) if none of the above conditions are true.  
  
The code that solves this problem is as extensible and flexible as possible.  
This is achieved through single responsibility rule and builder pattern.  
Using single responsibility modularity was achieved in:  
 - the definition of rules for FizzBuzz;  
 - iterations definition from 1..n;  
 - output format.
  
Using builder pattern a specific FizzBuzzer can be constructed for developer needs.  
  
Usage:  
-n flag specifies the n value.  
Example:  
```go run task1solution.go -n=15```
