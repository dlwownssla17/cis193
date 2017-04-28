package indexer

import (
	"coderive/src/tokenizer"
	"html"
	"strings"
)

// Match represents the match of the file for some query.
type Match struct {
	MatchLines []int
	FileInfo   *FileInfo
}

func convertForHTML(linesData []string) []string {
	linesDataForHTML := make([]string, len(linesData))
	for _, line := range linesData {
		lineForHTML := html.UnescapeString("&nbsp;")
		if line != "" {
			lineForHTML = strings.Replace(line, "\t", html.UnescapeString("&emsp;"), -1)
		}
		linesDataForHTML = append(linesDataForHTML, lineForHTML)
	}
	return linesDataForHTML
}

func findMatchesFromQMap(qMap *tokenizer.QueryMap) []*Match {
	if qMap == nil {
		return nil
	}

	queriesTextSearch := FindQueryTextSearches(qMap)

	matches := make([]*Match, 0)
	for _, q := range queriesTextSearch {
		q.FileInfo.LinesData = convertForHTML(q.FileInfo.LinesData)

		matches = append(matches, &Match{
			MatchLines: []int{-1}, // TODO: implement this
			FileInfo:   q.FileInfo,
		})
	}

	return matches
}

// FindMatches finds all matching files given the user's query.
func FindMatches(q string) []*Match {
	formattedQ := strings.Join(strings.Fields(q), " ")
	qMap := tokenizer.BuildQueryMap(formattedQ)
	return findMatchesFromQMap(qMap)
}
