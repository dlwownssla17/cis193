// Homework 8: CLI and Regex
// Due April 4, 2017 at 11:59pm
package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

// Problem 1: CLI
// Write a command line interface that prints out sequences of numbers.
//
// Usage of hw8:
// 	hw8 [flags] # prints out the sequence of numbers, each on a new line
// Flags:
//   -start int
//     	starting integer for the sequence (default 0)
//   -end   int
//      ending integer for the sequence, not inclusive (default 0)
//   -step  int
//      amount to skip by in each iteration (default 1)
//
// For example, executing `./hw8 -start=2 -end=5` should print out:
// 2
// 3
// 4
//
// Executing `./hw8 -start=2 -end=7 -step=3` should print out:
// 2
// 5
//
// Executing `./hw8 -start=10 -end=7 -step=-1` should print out:
// 10
// 9
// 8
//
// If the parameters are invalid (eg: positive step and start > end or
// negative step and start < end or invalid parameter values passed in),
// print out an error message using `log.Print(ln|f)?`.
//
// Feel free to do this section directly in the main() function.

func main() {
	printSequence()

	// s1 := `Sample text for testing: abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789 _+-.,!@#$%^&*();\/|<>"' 12345 -98.7 3.141 .6180 9,000 +42 555.123.4567  +1-(800)-555-2468 foo@demo.net  bar.ba@test.co.uk www.demo.com  http://foo.co.uk/ http://regexr.com/foo.html?q=bar https://mediatemple.net  asd123_.fg@adfdfj.sddfdsij.edu`
	// emails := GetEmails(s1)
	// fmt.Printf("Emails: %v\n", emails)

	// s2 := `123-456-7890 (123)-456-7890 123.456.7890 1234567890 123 123 456 7890 456 7890`
	// phoneNumbers := GetPhoneNumbers(s2)
	// fmt.Printf("Phone Numbers: %v\n", phoneNumbers)
}

func printSequence() {
	start := flag.Int("start", 0, "starting integer for the sequence (default 0)")
	end := flag.Int("end", 0, "ending integer for the sequence, not inclusive (default 0)")
	step := flag.Int("step", 1, "amount to skip by in each iteration (default 1)")

	flag.Parse()

	if len(flag.Args()) > 0 {
		log.Printf("invalid parameters: %v\n", flag.Args())
		return
	} else if (*step > 0 && *start > *end) || (*step < 0 && *start < *end) {
		log.Printf("invalid sequence (start: %v, end: %v, step: %v)\n", *start, *end, *step)
		return
	} else if *step == 0 {
		log.Println("step cannot be 0 (infinite sequence)")
		return
	}

	for i := *start; (*step > 0 && i < *end) || (*step < 0 && i > *end); i += *step {
		fmt.Println(i)
	}
}

// GetEmails takes in string input and returns a string slice of the
// emails found in the input string.
//
// Use regexp to extract all of the emails from the input string.
// Each email consists of the email name + "@" + domain + "." + top level domain.
// The email name should consist of only letters, numbers, underscores and dots.
// The domain should consist of only letters or dots.
// The top level domain must be "com", "org", "net" or "edu".
// between the domain and tld.
//
// You can assume that all email addresses will be surrounded by whitespace.
func GetEmails(s string) []string {
	re := regexp.MustCompile(`(\w|_|\.)+@[a-zA-Z\.]+\.(com|org|net|edu)`)
	return re.FindAllString(s, -1)
}

// GetPhoneNumbers takes in string input and returns a string slice of the
// phone numbers found in the input string.
//
// Use regexp to extract all of the phone numbers from the input string.
// Here are the formats phone numbers can be in for this problem:
// 215-555-3232
// (215)-555-3232
// 215.555.3232
// 2155553232
// 215 555 3232
//
// For your output, you should return a string slice of phone numbers with
// just the numbers (eg: "2158887744")
//
// You can assume that all phone numbers will be surrounded by whitespace.
func GetPhoneNumbers(s string) []string {
	re := regexp.MustCompile(`\d{3}-\d{3}-\d{4}|\(\d{3}\)-\d{3}-\d{4}|\d{3}\.\d{3}\.\d{4}|\d{10}|\d{3}\s\d{3}\s\d{4}`)
	return re.FindAllString(s, -1)
}
