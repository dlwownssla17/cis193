// Homework 5: Goroutines
// Due March 3, 2017 at 11:59pm
package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	// Feel free to use the main function for testing your functions
	hello := make(chan string, 5)
	hello <- "Hello world"
	hello <- "Привет мир"
	hello <- "Привіт Світ"
	hello <- "Witaj świecie"
	close(hello)
	fmt.Println(<-hello)
	for greeting := range hello {
		fmt.Println(greeting)
	}

	fmt.Println("---My Print Statements---")

	c := make(chan int, 5)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		c <- 5
		close(c)
	}()

	filtered := Filter(c, func(n int) bool { return n%2 == 1 })
	for v := range filtered {
		fmt.Println(v)
	}

	fmt.Println("---")

	tasks := []func() (string, error){
		func() (string, error) {
			time.Sleep(time.Second)
			fmt.Println(1)
			return "hello", nil
		},
		func() (string, error) {
			time.Sleep(time.Second)
			fmt.Println(2)
			return "world", nil
		},
		func() (string, error) {
			time.Sleep(time.Second)
			fmt.Println(3)
			return "cheese", nil
		},
		func() (string, error) {
			time.Sleep(time.Second)
			fmt.Println(4)
			return "burger", nil
		},
	}

	results := ConcurrentRetry(tasks, 2, 2)
	for result := range results {
		fmt.Println(result)
	}
}

// Filter copies values from the input channel into an output channel that match the filter function p
// The function p determines whether an int from the input channel c is sent on the output channel
func Filter(c <-chan int, p func(int) bool) <-chan int {
	filtered := make(chan int)
	go func() {
		for v := range c {
			if p(v) {
				filtered <- v
			}
		}
		close(filtered)
	}()
	return filtered
}

// Result is a type representing a single result with its index from a slice
type Result struct {
	index  int
	result string
}

// ConcurrentRetry runs all the tasks concurrently and sends the output in a Result channel
//
// concurrent is the limit on the number of tasks running in parallel. Your
// solution must not run more than `concurrent` number of tasks in parallel.
//
// retry is the number of times that the task should be attempted. If a task
// returns an error, the function should be retried immediately up to `retry`
// times. Only send the results of a task into the output channel if it does not error.
//
// Multiple instances of ConcurrentRetry should be able to run simultaneously
// without interfering with one another, so global variables should not be used.
// The function must return the channel without waiting for the tasks to
// execute, and all results should be sent on the output channel. Once all tasks
// have been completed, close the channel.
func ConcurrentRetry(tasks []func() (string, error), concurrent int, retry int) <-chan Result {
	results := make(chan Result, concurrent)
	var wg sync.WaitGroup

	for idx, task := range tasks {
		wg.Add(1)

		go func(idx int, task func() (string, error)) {
			defer wg.Done()

			for i := 0; i < retry; i++ {
				s, err := task()
				if err == nil {
					results <- Result{idx, s}
					break
				}
			}

		}(idx, task)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

// Task is an interface for types that process integers
type Task interface {
	Execute(int) (int, error)
}

func executeTasks(input int, tasks ...Task) <-chan struct {
	int
	error
} {
	results := make(chan struct {
		int
		error
	}, len(tasks))
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)

		go func(input int, task Task) {
			defer wg.Done()

			output, err := task.Execute(input)
			results <- struct {
				int
				error
			}{output, err}
		}(input, task)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

// Fastest returns the result of the fastest running task
// Fastest accepts any number of Task structs. If no tasks are submitted to
// Fastest(), it should return an error.
// You should return the result of a Task even if it errors.
// Do not leave any pending goroutines. Make sure all goroutines are cleaned up
// properly and any synchronizing mechanisms closed.
func Fastest(input int, tasks ...Task) (int, error) {
	if len(tasks) == 0 {
		errorString := "ERROR: no task submitted."
		log.Println(errorString)
		return 0, errors.New(errorString)
	}

	results := executeTasks(input, tasks...)

	fastest := <-results
	return fastest.int, fastest.error
}

// MapReduce takes any number of tasks, and feeds their results through reduce
// If no tasks are supplied, return an error.
// If any of the tasks error during their execution, return an error immediately.
// Once all tasks have completed successfully, return the value of reduce on
// their results in any order.
// Do not leave any pending goroutines. Make sure all goroutines are cleaned up
// properly and any synchronizing mechanisms closed.
func MapReduce(input int, reduce func(results []int) int, tasks ...Task) (int, error) {
	if len(tasks) == 0 {
		errorString := "ERROR: no task submitted."
		log.Println(errorString)
		return 0, errors.New(errorString)
	}

	results := executeTasks(input, tasks...)

	var resultsSlice []int
	for result := range results {
		if result.error != nil {
			return result.int, result.error
		}
		resultsSlice = append(resultsSlice, result.int)
	}

	return reduce(resultsSlice), nil
}
