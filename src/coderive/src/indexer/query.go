package indexer

import (
	"fmt"
	"strings"
)

func addTabs(lines []string) []string {
	for idx, line := range lines {
		lines[idx] = fmt.Sprintf("\t\t%s", line)
	}
	return lines
}

/* * */

// FileInfo is a set of key information for each file.
type FileInfo struct {
	Username string
	RepositoryName string
	BranchName string
	FilePath string
	Name string
	Link string
	NumLines int
	Data string
}

func (fileInfo *FileInfo) String() string {
	lines := strings.Split(fileInfo.Data, "\n")
	dataString := strings.Join(addTabs(lines), "\n")
	return fmt.Sprintf("FileInfo {\n" +
		"\tUsername: %s\n" +
		"\tRepositoryName: %s\n" +
		"\tBranchName: %s\n" +
		"\tFilePath: %s\n" +
		"\tName: %s\n" +
		"\tLink: %s\n" +
		"\tNumLines: %d\n" +
		"\tData:\n" +
		"\t\"\n" +
		"%s\n" +
		"\t\"\n" +
		"}", fileInfo.Username, fileInfo.RepositoryName, fileInfo.BranchName, fileInfo.FilePath,
		fileInfo.Name, fileInfo.Link, fileInfo.NumLines, dataString)
}

/* * */

type QueryTextSearch struct {
	FileInfo *FileInfo
}

func (q *QueryTextSearch) String() string {
	fileInfoString := q.FileInfo.String()
	lines := strings.Split(fileInfoString, "\n")
	fileInfoString = strings.Join(addTabs(lines), "\n")

	return fmt.Sprintf("QueryTextSearch {\n" +
		"\tFileInfo:\n" +
		"%s\n" +
		"}", fileInfoString)
}

/* * */

type QueryTextWordMatch struct {
	Count int
	Keywords []string
	MatchFiles []*struct {
		MatchLines []int
		FileInfo *FileInfo
	}
}

func (q *QueryTextWordMatch) String() string {
	matchFileStrings := make([]string, 0)
	for _, matchFile := range q.MatchFiles {
		fileInfoString := matchFile.FileInfo.String()
		lines := strings.Split(fileInfoString, "\n")
		fileInfoString = strings.Join(addTabs(lines), "\n")
		matchFileString := fmt.Sprintf("MatchFile {\n" +
			"\tMatchLines: %v\n" +
			"\tFileInfo:\n" +
			"%s\n" +
			"}", matchFile.MatchLines, fileInfoString)
		matchFileStrings = append(matchFileStrings, matchFileString)
	}
	matchFilesString := strings.Join(addTabs(matchFileStrings), "\n")

	return fmt.Sprintf("QueryTextWordMatch {\n" +
		"\tCount: %d\n" +
		"\tKeywords: %v\n" +
		"\tMatchFiles:\n" +
		"%s\n" +
		"}", q.Count, q.Keywords, matchFilesString)
}

/* * */

type QuerySemVarType struct {
	Type string
	Global bool
	MatchFiles []*struct {
		MatchLines []int
		FileInfo *FileInfo
	}
}
