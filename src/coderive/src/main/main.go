package main

import (
	"os"
	"coderive/src/crawler"
	"fmt"
	"coderive/src/indexer"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 3 && len(args) % 2 == 1 && args[0] == "crawl" {
		for i := 0; i < len(args) - 1; i += 2 {
			go func(username, repositoryName string) {
				repo := crawler.CrawlRepository(username, repositoryName)
				fmt.Printf("Finished crawling repository: %s\n", repo.Name)
			}(args[i], args[i+1])
		}
		return
	} else if len(args) == 2 && args[0] == "init" {
		if args[1] == "repositories" {
			crawler.DBRepositoriesInit()
			return
		} else if args[1] == "queries.textsearch" {
			indexer.DBQueriesTextSearchInit()
			return
		}
	} else if len(args) == 1 {
		if args[0] == "index" {
			indexer.IndexAll()
			return
		} else if args[0] == "server" {

			return
		} else if args[0] == "drop" {
			crawler.DBDrop()
			return
		}
	}

	fmt.Println("Usage:")
	fmt.Println("crawl [github username 1] [repository name 1] [github username 2] [repository name 2] ...")
	fmt.Println("---OR---")
	fmt.Println("index")
	fmt.Println("---OR---")
	fmt.Println("server")
	fmt.Println("---OR---")
	fmt.Println("init repositories")
	fmt.Println("---OR---")
	fmt.Println("init queries.textsearch")
	fmt.Println("---OR---")
	fmt.Println("drop")
}