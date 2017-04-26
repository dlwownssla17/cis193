package main

import (
	"os"
	"coderive/src/crawler"
	"fmt"
	"coderive/src/indexer"
	"coderive/src/common"
)

func usage() {
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

func main() {
	args := os.Args[1:]

	if len(args) >= 3 && len(args) % 2 == 1 && args[0] == "crawl" {
		fmt.Println("Started crawling specified repositories")

		for i := 0; i < len(args) - 1; i += 2 {
			go func(username, repositoryName string) {
				repo := crawler.CrawlRepository(username, repositoryName)

				if &repo != nil {
					fmt.Printf("Finished crawling repository: %s\n", repositoryName)
				} else {
					fmt.Printf("Repository already exists: %s\n", repositoryName)
				}
			}(args[i], args[i+1])
		}
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
		} else if args[0] == "drop" {
			fmt.Println("Started dropping entire database")

			common.DBDrop()

			fmt.Println("Finished dropping entire database")
			return
		}
	}

	usage()
}