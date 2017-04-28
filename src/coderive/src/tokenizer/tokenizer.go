package tokenizer

import (
	"strconv"
	"strings"
)

// QueryMap represents the user's query in a structure appropriate for database lookup.
type QueryMap map[string]interface{}

var delims = map[byte]bool{
	'[': false,
	']': false,
	';': false,
	',': false,
}

var linesKeys = map[string]bool{
	"eq":  false,
	"lt":  false,
	"gt":  false,
	"lte": false,
	"gte": false,
	"ne":  false,
}

var textKeys = map[string]bool{
	"val":   false,
	"regex": false,
}

func tokenize(q string) []string {
	tokens := make([]string, 0)
	i := 1

	q = strings.TrimSpace(q)
	for q != "" {
		if q[0] == '"' || q[0] == '\'' { // text value (including quotes)
			for i < len(q) &&
				!((q[0] == '"' && q[i] == '"') || (q[0] == '\'' && q[i] == '\'') && q[i-1] != '\\') {
				i++
			}
			if i < len(q) {
				i++
			}
		} else if _, ok := delims[q[0]]; !ok { // non-delimiter
			isDelim := false
			for i < len(q) && !isDelim {
				if _, isDelim = delims[q[i]]; !isDelim {
					i++
				}
			}
		}

		tokens = append(tokens, strings.TrimSpace(q[:i]))
		q = strings.TrimSpace(q[i:])
		i = 1
	}

	return tokens
}

func buildQueryMapHelper(tokens []string) *QueryMap {
	qMap := make(QueryMap)
	i := 0

	for i < len(tokens) {
		qType := tokens[i]
		i++
		switch qType {
		case ";":
			continue
		case "lines":
			linesMap := make(map[string]int)

			if i >= len(tokens) || tokens[i] != "[" {
				return nil
			}
			i++

			if i >= len(tokens) {
				return nil
			}
			if _, ok := linesKeys[tokens[i]]; !ok {
				return nil
			}
			linesOp := tokens[i]
			i++

			if i >= len(tokens) || tokens[i] != "[" {
				return nil
			}
			i++

			if i >= len(tokens) {
				return nil
			}
			linesThreshold, err := strconv.Atoi(tokens[i])
			if err != nil {
				return nil
			}
			i++

			if i >= len(tokens) || tokens[i] != "]" {
				return nil
			}
			i++

			if i >= len(tokens) || tokens[i] != "]" {
				return nil
			}
			i++

			linesMap[linesOp] = linesThreshold
			qMap[qType] = linesMap
		case "text":
			textMap := make(map[string]interface{})

			if i >= len(tokens) || tokens[i] != "[" {
				return nil
			}
			i++

			continueProcessing := true
			for continueProcessing {
				if i >= len(tokens) {
					return nil
				}
				if _, ok := textKeys[tokens[i]]; !ok {
					return nil
				}
				if _, ok := textMap[tokens[i]]; ok {
					return nil
				}
				textKey := tokens[i]
				i++

				if i >= len(tokens) || tokens[i] != "[" {
					return nil
				}
				i++

				if i >= len(tokens) {
					return nil
				}
				switch textKey {
				case "val":
					if len(tokens[i]) < 2 ||
						!(tokens[i][0] == '"' && tokens[i][len(tokens[i])-1] == '"') &&
							!(tokens[i][0] == '\'' && tokens[i][len(tokens[i])-1] == '\'') {
						return nil
					}

					textMap[textKey] = tokens[i][1 : len(tokens[i])-1]
				case "regex":
					isRegex, err := strconv.ParseBool(tokens[i])
					if err != nil {
						return nil
					}

					textMap[textKey] = isRegex
				}
				i++

				if i >= len(tokens) || tokens[i] != "]" {
					return nil
				}
				i++

				if i >= len(tokens) || !(tokens[i] == "]" || tokens[i] == ",") {
					return nil
				} else if tokens[i] == "]" {
					continueProcessing = false
				}
				i++
			}

			if _, ok := textMap["val"]; !ok {
				return nil
			}
			if _, ok := textMap["regex"]; !ok {
				textMap["regex"] = false
			}
			qMap[qType] = textMap
		default:
			return nil
		}
	}

	return &qMap
}

// BuildQueryMap builds a QueryMap given the user's query string.
func BuildQueryMap(q string) *QueryMap {
	tokens := tokenize(q)
	return buildQueryMapHelper(tokens)
}
