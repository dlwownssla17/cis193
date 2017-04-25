package crawler

import (
	"fmt"
	"strings"
)

func addTabs(lines []string) []string {
	for idx, line := range(lines) {
		lines[idx] = fmt.Sprintf("\t\t%s", line)
	}
	return lines
}

/* * */

// File is a representation of a go source code file.
type File struct {
	Name string
	Link string
	NumLines int
	Data string
}

func (f *File) String() string {
	lines := strings.Split(f.Data, "\n")
	dataString := strings.Join(addTabs(lines), "\n")
	return fmt.Sprintf("File {\n\tName: %s\n\tLink: %s\n\tNumLines: %d\n\tData:\n\t\"\n%s\n\t\"\n}", f.Name, f.Link, f.NumLines, dataString)
}

/* * */

// Directory contains go source code files and subdirectories.
type Directory struct {
	Name           string
	Files          []*File
	Subdirectories []*Directory
}

func (dir *Directory) GetNumFiles() int {
	sum := 0
	for _, subdirectory := range(dir.Subdirectories) {
		sum += subdirectory.GetNumFiles()
	}
	return len(dir.Files) + sum
}

func (dir *Directory) String() string {
	fileStrings := make([]string, 0)
	for _, file := range(dir.Files) {
		fileString := file.String()
		lines := strings.Split(fileString, "\n")
		fileStrings = append(fileStrings, strings.Join(addTabs(lines), "\n"))
	}
	filesString := strings.Join(fileStrings, ",\n")

	subdirectoryStrings := make([]string, 0)
	for _, subdirectory := range(dir.Subdirectories) {
		subdirectoryString := subdirectory.String()
		lines := strings.Split(subdirectoryString, "\n")
		subdirectoryStrings = append(subdirectoryStrings, strings.Join(addTabs(lines), "\n"))
	}
	subdirectoriesString := strings.Join(subdirectoryStrings, ",\n")

	return fmt.Sprintf("Directory {\n\tName: %s\n\tFiles:\n%s\n\tSubdirectories:\n%s\n}", dir.Name, filesString, subdirectoriesString)
}

/* * */

// Branch is a representation of a branch within a user's GitHub repository.
type Branch struct {
	Name string
	Root *Directory
}

func (br *Branch) GetNumFiles() int {
	return br.Root.GetNumFiles()
}

func (br *Branch) String() string {
	rootString := br.Root.String()
	lines := strings.Split(rootString, "\n")
	rootString = strings.Join(addTabs(lines), "\n")

	return fmt.Sprintf("Branch {\n\tName: %s\n\tRoot:\n%s\n}", br.Name, rootString)
}

/* * */

// Repository is a representation of a user's GitHub repository.
type Repository struct {
	Username string
	Name string
	Branches []*Branch
	Processed bool
}

func (repo *Repository) GetNumFiles() int {
	sum := 0
	for _, branch := range(repo.Branches) {
		sum += branch.GetNumFiles()
	}
	return sum
}

func (repo *Repository) String() string {
	branchStrings := make([]string, 0)
	for _, branch := range(repo.Branches) {
		branchString := branch.String()
		lines := strings.Split(branchString, "\n")
		branchStrings = append(branchStrings, strings.Join(addTabs(lines), "\n"))
	}
	branchesString := strings.Join(branchStrings, ",\n")

	return fmt.Sprintf("Repository {\n\tUsername: %s\n\tName: %s\n\tBranches:\n%s\n\tProcessed: %v\n}", repo.Username, repo.Name, branchesString, repo.Processed)
}