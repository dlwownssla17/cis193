// Homework 4: Concurrency
// Due February 21, 2017 at 11:59pm
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// Feel free to use the main function for testing your functions
	hello := map[string]string{
		"こんにちは": "世界",
		"你好":    "世界",
		"안녕하세요": "세계",
	}
	for k, v := range hello {
		fmt.Printf("%s, %s\n", strings.Title(k), v)
	}

	fmt.Println("---My Tests---")

	FileSum("hw4_test1", "hw4_result1")
	FileSum("hw4_test1", "hw4_result2")
	FileSum("hw4_test2", "hw4_result2")

	dir := PennDirectory{}
	dir.directory = make(map[int]string)
	fmt.Println(dir.Get(12345678))
	dir.Add(12345678, "jjlee")
	dir.Add(87654321, "adelq")
	dir.Add(12345678, "prakharb")
	fmt.Println(dir.Get(12345678))
	fmt.Println(dir.Get(87654321))
	dir.Remove(12345678)
	fmt.Println(dir.Get(12345678))
}

func checkErr(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}

// Problem 1a: File processing
// You will be provided an input file consisting of integers, one on each line.
// Your task is to read the input file, sum all the integers, and write the
// result to a separate file.

// FileSum sums the integers in input and writes them to an output file.
// The two parameters, input and output, are the filenames of those files.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func FileSum(input, output string) {
	fi, err := os.Open(input)
	checkErr(err)
	fo, err := os.Create(output)
	checkErr(err)

	defer fi.Close()
	defer fo.Close()

	IOSum(fi, fo)
}

// Problem 1b: IO processing with interfaces
// You must do the exact same task as above, but instead of being passed 2
// filenames, you are passed 2 interfaces: io.Reader and io.Writer.
// See https://golang.org/pkg/io/ for information about these two interfaces.
// Note that os.Open returns an io.Reader, and os.Create returns an io.Writer.

// IOSum sums the integers in input and writes them to output
// The two parameters, input and output, are interfaces for io.Reader and
// io.Writer. The type signatures for these interfaces is in the Go
// documentation.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func IOSum(input io.Reader, output io.Writer) {
	sc := bufio.NewScanner(input)
	w := bufio.NewWriter(output)

	sum := 0
	for sc.Scan() {
		n, err := strconv.Atoi(sc.Text())
		checkErr(err)
		sum += n
	}
	checkErr(sc.Err())

	_, err := w.WriteString(fmt.Sprintf("%d\n", sum))
	checkErr(err)

	checkErr(w.Flush())
}

// Problem 2: Concurrent map access
// Maps in Go [are not safe for concurrent use](https://golang.org/doc/faq#atomic_maps).
// For this assignment, you will be building a custom map type that allows for
// concurrent access to the map using mutexes.
// The map is expected to have concurrent readers but only 1 writer can have
// access to the map.

// PennDirectory is a mapping from PennID number to PennKey (12345678 -> adelq).
// You may only add *private* fields to this struct.
// Hint: Use an embedded sync.RWMutex, see lecture 2 for a review on embedding
type PennDirectory struct {
	directory map[int]string
	mu        sync.RWMutex
}

// Add inserts a new student to the Penn Directory.
// Add should obtain a write lock, and should not allow any concurrent reads or
// writes to the map.
// You may NOT write over existing data - simply raise a warning.
func (d *PennDirectory) Add(id int, name string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if val, ok := d.directory[id]; ok {
		log.Println(fmt.Sprintf("WARNING: %d -> %s already exists.", id, val))
		return
	}
	d.directory[id] = name
}

// Get fetches a student from the Penn Directory by their PennID.
// Get should obtain a read lock, and should allow concurrent read access but
// not write access.
func (d *PennDirectory) Get(id int) string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.directory[id]
}

// Remove deletes a student to the Penn Directory.
// Remove should obtain a write lock, and should not allow any concurrent reads
// or writes to the map.
func (d *PennDirectory) Remove(id int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.directory, id)
}
