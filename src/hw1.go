// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, دنيا!")
	fmt.Println(ParsePhone("123-456-7890"))
	fmt.Println(ParsePhone("1 2 3 4 5 6 7 8 9 0"))
	fmt.Println(Anagram("aba", "BaA"))
	fmt.Println(Anagram("aaa", "aa"))
	fmt.Println(Anagram("abb", "baa"))
	fmt.Println(FindEvens([]int{-3, -2, -1, 0, 1, 2, 3, 4}))
	fmt.Println(SliceProduct([]int{}))
	fmt.Println(SliceProduct([]int{1, -2, 3}))
	fmt.Println(Unique([]int{1, 1, 0, 1, 3, 1, 0, 0, 0, 2, 3, 2}))
	fmt.Println(InvertMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 3}))
	fmt.Println(TopCharacters("aaabbbbccddddd", 3))
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	res := [14]rune{0: '(', 4: ')', 5: ' ', 9: '-'}
	for i, resPtr := 0, 0; resPtr < len(res); {
		if resPtr == 0 || resPtr == 4 || resPtr == 5 || resPtr == 9 {
			resPtr++
		} else {
			if r := rune(phone[i]); unicode.IsDigit(r) {
				res[resPtr] = r
				resPtr++
			}
			i++
		}
	}
	return string(res[:])
}

// Helper function that takes a string and builds a map[rune]int
// mapping from the characters that occur in the string to their
// occurrences
func BuildCount(s string) map[rune]int {
	count := make(map[rune]int)
	for _, c := range s {
		count[rune(c)]++
	}
	return count
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	count := BuildCount(strings.ToLower(s1))
	for _, c := range strings.ToLower(s2) {
		r := rune(c)
		if count[r] == 0 {
			return false
		}
		count[r]--
	}
	return true
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	var evens []int
	for _, n := range e {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	return evens
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	if len(e) == 0 {
		return 0
	}

	product := 1
	for _, n := range e {
		product *= n
	}
	return product
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	var x struct{}
	var unique []int
	exists := make(map[int]struct{})
	for _, n := range e {
		if _, ok := exists[n]; !ok {
			unique = append(unique, n)
			exists[n] = x
		}
	}
	return unique
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
func InvertMap(kv map[string]int) map[int]string {
	vk := make(map[int]string)
	for k, v := range kv {
		vk[v] = k
	}
	return vk
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	top := make(map[rune]int)
	count := BuildCount(s)
	for r, n := range count {
		if n > k {
			top[r] = n
		}
	}
	return top
}
