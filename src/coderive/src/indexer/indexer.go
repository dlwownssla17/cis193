package indexer

import (
	"coderive/src/crawler"
	"fmt"
	"sync"
	"strings"
)

func buildQueriesThroughDirectory(qs []*QueryTextSearch, dir *crawler.Directory,
	username, repositoryName, branchName, currentFilePath string) []*QueryTextSearch {
	for _, file := range dir.Files {
		fileInfo := &FileInfo{
			Username:       username,
			RepositoryName: repositoryName,
			BranchName:     branchName,
			FilePath:       fmt.Sprintf("%s%s", currentFilePath, file.Name),
			Name:           file.Name,
			Link:           file.Link,
			NumLines:       file.NumLines,
			LinesData:      strings.Split(file.Data, "\n"),
			FormattedData:  strings.Join(strings.Fields(file.Data), " "),
		}
		q := &QueryTextSearch{
			FileInfo: fileInfo,
		}
		qs = append(qs, q)
	}

	for _, subdirectory := range dir.Subdirectories {
		qs = buildQueriesThroughDirectory(qs, subdirectory,
			username, repositoryName, branchName, fmt.Sprintf("%s%s/", currentFilePath, subdirectory.Name))
	}

	return qs
}

func toQueriesTextSearch(repo crawler.Repository) []*QueryTextSearch {
	qs := make([]*QueryTextSearch, 0)
	for _, branch := range repo.Branches {
		qs = buildQueriesThroughDirectory(qs, branch.Root, repo.Username, repo.Name, branch.Name, "/")
	}
	return qs
}

/* * */

// IndexAll processes all the crawled documents in the repositories collection to update all the queries collections.
func IndexAll() int {
	var wg sync.WaitGroup

	reposToProcess := GetAllRepositoriesToProcess()

	for _, repo := range reposToProcess {
		wg.Add(3)

		go func(repo crawler.Repository) {
			defer wg.Done()

			qs := toQueriesTextSearch(repo)
			for _, q := range qs {
				SaveQueryTextSearch(q)
			}
		}(repo)

		go func(repo crawler.Repository) {
			defer wg.Done()


		}(repo)

		go func(repo crawler.Repository) {
			defer wg.Done()


		}(repo)
	}

	wg.Wait()

	UpdateAllRepositoriesProcessed()

	return len(reposToProcess)
}