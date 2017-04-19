package main

import (
	"fmt"
	"coderive/src/crawler"
)

func main() {
	//repo := GetRepository("dlwownssla17", "cis193")
	repo := crawler.GetRepository("yakumioto", "CrawlerIShadowsocks")
	fmt.Println(&repo)
	fmt.Printf("Num Files: %d\n", repo.GetNumFiles())
	fmt.Printf("Num Branches: %d\n", len(repo.Branches))
}