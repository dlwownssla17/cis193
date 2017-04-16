// Homework 0: Hello Go!
// Due January 24, 2017 at 11:59pm
package main

import "fmt"
import "math"

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, दुनिया!")
	fmt.Println(IsPalindrome("aba"))
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return ""
	}
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	if n == 1 || n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n&i == 0 {
			return false
		}
	}
	return true
}

// IsPalindromeHW3 checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
