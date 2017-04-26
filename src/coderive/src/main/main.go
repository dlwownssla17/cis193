package main

import (
	"os"
	"coderive/src/crawler"
	"fmt"
	"coderive/src/indexer"
	"coderive/src/common"
	"sync"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("crawl [github username 1] [repository name 1] [github username 2] [repository name 2] ...")
	fmt.Println("---OR---")
	fmt.Println("index")
	fmt.Println("---OR---")
	fmt.Println("server")
	fmt.Println("---OR---")
	fmt.Println("init [repositories|queries.textsearch|queries.textwordmatch|queries.semvartype]?")
	fmt.Println("---OR---")
	fmt.Println("drop")
}

func crawlAll(args []string) {
	var wg sync.WaitGroup

	for i := 1; i < len(args); i += 2 {
		wg.Add(1)

		go func(username, repositoryName string) {
			defer wg.Done()

			repo := crawler.CrawlRepository(username, repositoryName)

			if repo != nil {
				fmt.Printf("Done crawling repository: %s %s\n", username, repositoryName)
			} else {
				fmt.Printf("Repository already exists: %s %s\n", username, repositoryName)
			}
		}(args[i], args[i+1])
	}

	wg.Wait()
}

func main() {
	args := os.Args[1:]

	if len(args) >= 3 && len(args) % 2 == 1 && args[0] == "crawl" {
		fmt.Println("Started crawling specified repositories")

		crawlAll(args)

		fmt.Println("Finished crawling specified repositories")
		return
	} else if len(args) == 2 && args[0] == "init" &&
		(args[1] == "repositories" || args[1] == "queries.textsearch") {
		fmt.Printf("Started initializing %s db\n", args[1])

		common.DBCollectionInit(args[1])

		fmt.Printf("Finished initializing %s db\n", args[1])
		return
	} else if len(args) == 1 {
		if args[0] == "index" {
			fmt.Println("Started indexing")

			count := indexer.IndexAll()

			fmt.Printf("Finished indexing: %d repositories processed\n", count)
			return
		} else if args[0] == "server" {

			return
		} else if args[0] == "init" {
			fmt.Println("Started initializing all dbs")

			dbs := []string{"repositories", "queries.textsearch"}
			for _, db := range dbs {
				common.DBCollectionInit(db)
			}

			fmt.Println("Finished initializing all dbs")
			return
		} else if args[0] == "drop" {
			fmt.Println("Started dropping entire database")

			common.DBDrop()

			fmt.Println("Finished dropping entire database")
			return
		}
	}

	usage()
}