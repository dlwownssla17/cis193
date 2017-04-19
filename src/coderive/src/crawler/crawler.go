package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

/* * */

func getBranchNames(username, repositoryName string) []string {
	doc, err := goquery.NewDocument(GitHubRepositoryUrlifyWithParams(username, repositoryName))
	if err != nil {
		log.Fatal(err)
	}
	branchNamesDoc := doc.Find("div[data-tab-filter=branches] a")

	branchNames := make([]string, 0)
	for i := 0; i < branchNamesDoc.Length(); i++ {
		branchName, ok := branchNamesDoc.Eq(i).Attr("data-name")
		if !ok {
			log.Fatalf("GitHub User: %s\nRepository Name: %s\nBranch HTML element is formatted unexpectedly.\n", username, repositoryName)
		}
		branchNames = append(branchNames, branchName)
	}
	return branchNames
}

/* * */

func getFileWithParams(username, repositoryName, branchName, filePath, filename string) *File {
	doc, err := goquery.NewDocument(GitHubRawContentUrlifyWithParams(username, repositoryName, branchName, filePath))
	if err != nil {
		log.Fatal(err)
	}

	data := doc.Find("body").Eq(0).Text()
	numLines := len(strings.Split(data, "\n"))

	return &File{
		Name: filename,
		Link: GitHubUrlifyWithParams(username, repositoryName, branchName, filePath, true),
		NumLines: numLines,
		Data: data,
	}
}

/* * */

func isLinkToFile(username, repositoryName, partURL string) bool {
	if len(username) == 0 || len(repositoryName) == 0 {
		log.Fatalln("Invalid username or repository name.")
	}
	partURLSubstring := fmt.Sprintf("%s/%s/", username, repositoryName)
	idxPartURLSubstring := strings.Index(partURL, partURLSubstring)
	idxAfter := idxPartURLSubstring + len(partURLSubstring)
	idxBlob := idxAfter + strings.Index(partURL[idxAfter:], "blob")
	return idxPartURLSubstring != -1 && idxBlob == idxAfter
}

func hasGoExtension(name string) bool {
	return strings.HasSuffix(name, ".go")
}

func appendPath(originalPath, additionalPath string) string {
	if len(originalPath) == 0 {
		return additionalPath
	}
	if len(additionalPath) == 0 {
		return originalPath
	}

	if originalPath[len(originalPath) - 1] == '/' {
		originalPath = originalPath[:len(originalPath) - 1]
	}
	if additionalPath[0] == '/' {
		additionalPath = additionalPath[1:]
	}
	return fmt.Sprintf("%s/%s", originalPath, additionalPath)
}

func getDirectory(username, repositoryName, branchName, filePath, directoryName string) *Directory {
	doc, err := goquery.NewDocument(GitHubUrlifyWithParams(username, repositoryName, branchName, filePath, false))
	if err != nil {
		log.Fatal(err)
	}
	childrenDoc := doc.Find("table.files td.content a")

	files := make([]*File, 0)
	directories := make([]*Directory, 0)
	for i := 0; i < childrenDoc.Length(); i++ {
		childDoc := childrenDoc.Eq(i)
		childName := childDoc.Text()
		childLink, ok := childDoc.Attr("href")
		if !ok {
			log.Fatalf("GitHub User: %s\nRepository Name: %s\nBranch Name: %s\n File Path: %s\nHTML element for %s is missing a link unexpectedly.\n", username, repositoryName, branchName, filePath, childName)
		}

		extendedPath := appendPath(filePath, childName)
		isFile := isLinkToFile(username, repositoryName, childLink)
		if isFile && hasGoExtension(childName) {
			childFile := getFileWithParams(username, repositoryName, branchName, extendedPath, childName)
			files = append(files, childFile)
		} else if !isFile {
			childDirectory := getDirectory(username, repositoryName, branchName, extendedPath, childName)
			directories = append(directories, childDirectory)
		}
	}

	return &Directory{
		Name:           directoryName,
		Files:          files,
		Subdirectories: directories,
	}
}

/* * */

func getBranch(username, repositoryName, branchName string) *Branch {
	return &Branch{
		Name: branchName,
		Root: getDirectory(username, repositoryName, branchName, "", ""),
	}
}

/* * */

// GetRepository crawls through the specified repository and builds the corresponding Repository instance
func GetRepository(username, repositoryName string) Repository {
	branchNames := getBranchNames(username, repositoryName)

	branches := make([]*Branch, 0)
	for i := 0; i < len(branchNames); i++ {
		branches = append(branches, getBranch(username, repositoryName, branchNames[i]))
	}

	return Repository{
		Username: username,
		Name: repositoryName,
		Branches: branches,
	}
}