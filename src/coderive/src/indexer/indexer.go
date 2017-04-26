package indexer

import (
	"coderive/src/crawler"
	"fmt"
)

func buildQueriesThroughDirectory(qs []*QueryTextSearch, dir *crawler.Directory,
	username, repositoryName, branchName, currentFilePath string) []*QueryTextSearch {
	for _, file := range dir.Files {
		fileInfo := &FileInfo{
			Username: username,
			RepositoryName: repositoryName,
			BranchName: branchName,
			FilePath: fmt.Sprintf("%s%s", currentFilePath, file.Name),
			Name: file.Name,
			Link: file.Link,
			NumLines: file.NumLines,
			Data: file.Data,
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

// IndexAll processes all the crawled documents in the repositories collection to update all the queries collections.
func IndexAll() int {
	reposToProcess := GetAllRepositoriesToProcess()

	for _, repo := range reposToProcess {
		qs := toQueriesTextSearch(repo)
		for _, q := range qs {
			SaveQueryTextSearch(q)
		}
	}

	UpdateAllRepositoriesProcessed()

	return len(reposToProcess)
}