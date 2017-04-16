// Homework 3: Interfaces
// Due February 14, 2017 at 11:59pm
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Feel free to use the main function for testing your functions
	hello := map[string]string{
		"hello":   "world",
		"hola":    "mundo",
		"bonjour": "monde",
	}
	for k, v := range hello {
		fmt.Printf("%s, %s\n", strings.Title(k), v)
	}

	fmt.Println("---My Print Statements---")

	fmt.Println('Z' < 'a')
	ps := make(PersonSlice, 5)
	ps[0] = NewPerson("John", "Smith")
	ps[1] = NewPerson("John", "Snow")
	ps[2] = NewPerson("Adam", "Smith")
	ps[3] = NewPerson("John", "Smit")
	ps[4] = NewPerson("John", "Smith")
	sort.Sort(ps)
	for i := 0; i < len(ps); i++ {
		fmt.Println(ps[i])
	}

	psPal := make(PersonSlice, 3)
	psPal[0] = &Person{-1, "John", "Smith"}
	psPal[1] = &Person{-1, "Adam", "Smith"}
	psPal[2] = &Person{-1, "John", "Smith"}
	fmt.Println(IsPalindrome(psPal))

	fmt.Println(Fold([]int{1, 2, 4}, 0, add))
	fmt.Println(Fold([]int{}, 42, mul))
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

// Problem 1: Sorting Names
// Sorting in Go is done through interfaces!
// To sort a collection (such as a slice), the type must satisfy sort.Interface,
// which requires 3 methods: Len() int, Less(i, j int) bool, and Swap(i, j int).
// To actually sort a slice, you need to first implement all 3 methods on a
// custom type, and then call sort.Sort on your the PersonSlice type.
// See the Go documentation: https://golang.org/pkg/sort/ for full details.

// Person stores a simple profile. These should be sorted by alphabetical order
// by last name, followed by the first name, followed by the ID. You can assume
// the ID will be unique, but the names need not be unique.
// Sorting should be case-sensitive and UTF-8 aware.
type Person struct {
	ID        int
	FirstName string
	LastName  string
}

// PersonSlice is a slice of Persons.
type PersonSlice []*Person

func (ps PersonSlice) Len() int {
	return len(ps)
}

func compare(s1, s2 string) int {
	r1, r2 := []rune(s1), []rune(s2)

	min := len(r1)
	if min > len(r2) {
		min = len(r2)
	}

	for i := 0; i < min; i++ {
		if r1[i] != r2[i] {
			if r1[i] < r2[i] {
				return -1
			}
			return 1
		}
	}

	if len(r1) < len(r2) {
		return -1
	} else if len(r1) > len(r2) {
		return 1
	}
	return 0
}

func (ps PersonSlice) Less(i, j int) bool {
	if compareLastName := compare(ps[i].LastName, ps[j].LastName); compareLastName != 0 {
		return compareLastName == -1
	}
	if compareFirstName := compare(ps[i].FirstName, ps[j].FirstName); compareFirstName != 0 {
		return compareFirstName == -1
	}
	return ps[i].ID < ps[j].ID
}

func (ps PersonSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// IDCtr is the automatically assigned ID for each new Person
var IDCtr = 1

// NewPerson is a constructor for Person. ID should be assigned automatically in
// sequential order, starting at 1 for the first Person created.
func NewPerson(first, last string) *Person {
	p := &Person{IDCtr, first, last}
	IDCtr++
	return p
}

// Problem 2: IsPalindrome Redux
// Using a function that simply requires sort.Interface, you should be able to
// check if a sequence is a palindrome. You may use, adapt, or modify your code
// from HW0. Note that the input does not need to be a string: any type which
// satisfies sort.Interface can (and will) be used to test. This means that the
// only functionality you should use should come from the sort.Interface methods
// Ex: [1, 2, 1] => true

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		if s.Less(i, s.Len()-1-i) || s.Less(s.Len()-1-i, i) {
			return false
		}
	}
	return true
}

// Problem 3: Functional Programming
// Write a function Fold which applies a function repeatedly on a slice,
// producing a single value via repeated application of an input function.
// The behavior of Fold should be as follows:
//   - When s is empty, return v (default value)
//   - When s has 1 value (x0), apply f once: f(v, x0)
//   - When s has 2 values (x0, x1), apply f twice, from left to right: f(f(v, x0), x1)
//   - Continue this pattern recursively to obtain the final result.

// Fold applies a left to right application of f on s starting with v.
// Note the argument signature of f - func(int, int) int.
// This means f is a function which has 2 int arguments and returns an int.
func Fold(s []int, v int, f func(int, int) int) int {
	for i := 0; i < len(s); i++ {
		v = f(v, s[i])
	}
	return v
}
