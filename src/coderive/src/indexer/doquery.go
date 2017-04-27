package indexer

// Matches represents the match of the file for some query.
type Match struct {
	MatchLines []int
	FileInfo *FileInfo
}

