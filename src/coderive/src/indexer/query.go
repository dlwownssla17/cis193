package indexer

import (
	"coderive/src/common"
	"fmt"
	"strings"
)

// FileInfo is a set of key information for each file.
type FileInfo struct {
	Username       string
	RepositoryName string
	BranchName     string
	FilePath       string
	Name           string
	Link           string
	NumLines       int
	LinesData      []string
	FormattedData  string
}

func (fileInfo *FileInfo) String() string {
	dataString := strings.Join(common.AddTabs(fileInfo.LinesData), "\n")
	return fmt.Sprintf("FileInfo {\n"+
		"\tUsername: %s\n"+
		"\tRepositoryName: %s\n"+
		"\tBranchName: %s\n"+
		"\tFilePath: %s\n"+
		"\tName: %s\n"+
		"\tLink: %s\n"+
		"\tNumLines: %d\n"+
		"\tRawData:\n"+
		"\t\"\n"+
		"%s\n"+
		"\t\"\n"+
		"\tData: %s\n"+
		"}", fileInfo.Username, fileInfo.RepositoryName, fileInfo.BranchName, fileInfo.FilePath,
		fileInfo.Name, fileInfo.Link, fileInfo.NumLines, dataString, fileInfo.FormattedData)
}

/* * */

// QueryTextSearch represents query by conventional text search via suboptimal source code string analysis.
type QueryTextSearch struct {
	FileInfo *FileInfo
}

func (q *QueryTextSearch) String() string {
	fileInfoString := q.FileInfo.String()
	lines := strings.Split(fileInfoString, "\n")
	fileInfoString = strings.Join(common.AddTabs(lines), "\n")

	return fmt.Sprintf("QueryTextSearch {\n"+
		"\tFileInfo:\n"+
		"%s\n"+
		"}", fileInfoString)
}

/* * */

// QueryTextWordMatch represents query by conventional text search via efficient inverted indexing.
type QueryTextWordMatch struct {
	Count    int
	Keywords []string
	Matches  []*Match
}

func (q *QueryTextWordMatch) String() string {
	matchFileStrings := make([]string, 0)
	for _, matchFile := range q.Matches {
		fileInfoString := matchFile.FileInfo.String()
		lines := strings.Split(fileInfoString, "\n")
		fileInfoString = strings.Join(common.AddTabs(lines), "\n")
		matchFileString := fmt.Sprintf("MatchFile {\n"+
			"\tMatchLines: %v\n"+
			"\tFileInfo:\n"+
			"%s\n"+
			"}", matchFile.MatchLines, fileInfoString)
		matchFileStrings = append(matchFileStrings, matchFileString)
	}
	matchFilesString := strings.Join(common.AddTabs(matchFileStrings), "\n")

	return fmt.Sprintf("QueryTextWordMatch {\n"+
		"\tCount: %d\n"+
		"\tKeywords: %v\n"+
		"\tMatchFiles:\n"+
		"%s\n"+
		"}", q.Count, q.Keywords, matchFilesString)
}

/* * */

// QuerySemVarType represents query by semantics variable type and globality.
type QuerySemVarType struct {
	Type    string
	Global  bool
	Matches []*Match
}

func (q *QuerySemVarType) String() string {
	matchFileStrings := make([]string, 0)
	for _, matchFile := range q.Matches {
		fileInfoString := matchFile.FileInfo.String()
		lines := strings.Split(fileInfoString, "\n")
		fileInfoString = strings.Join(common.AddTabs(lines), "\n")
		matchFileString := fmt.Sprintf("MatchFile {\n"+
			"\tMatchLines: %v\n"+
			"\tFileInfo:\n"+
			"%s\n"+
			"}", matchFile.MatchLines, fileInfoString)
		matchFileStrings = append(matchFileStrings, matchFileString)
	}
	matchFilesString := strings.Join(common.AddTabs(matchFileStrings), "\n")

	return fmt.Sprintf("QuerySemVarType {\n"+
		"\tType: %s\n"+
		"\tGlobal: %v\n"+
		"\tMatchFiles:\n"+
		"%s\n"+
		"}", q.Type, q.Global, matchFilesString)
}
