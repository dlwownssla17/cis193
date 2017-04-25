package main

import (
	"os"
	"coderive/src/crawler"
	"fmt"
)

func main() {
	args := os.Args[1:]
	if len(args) == 3 && args[0] == "crawl" {
		repo := crawler.CrawlRepository(args[1], args[2])
		fmt.Println(&repo)
		return
	} else if len(args) == 1 && args[0] == "index" {

		return
	} else if len(args) == 1 && args[0] == "server" {

		return
	}

	fmt.Println("Usage:")
	fmt.Println("crawl [github username] [repository name]")
	fmt.Println("---OR---")
	fmt.Println("index")
	fmt.Println("---OR---")
	fmt.Println("server")
}