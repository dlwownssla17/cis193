package tokenizer

import (
	"strings"
	"strconv"
	"fmt"
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
	"eq": false,
	"lt": false,
	"gt": false,
	"lte": false,
	"gte": false,
	"ne": false,
}

var textKeys = map[string]bool{
	"val": false,
	"regex": false,
}

func tokenize(q string) []string {
	tokens := make([]string, 0)
	i := 1

	q = strings.TrimSpace(q)
	for q != "" {
		if q[0] == '"' { // text value (including quotes)
			for i < len(q) && !(q[i] == '"' && q[i-1] != '\\') {
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

// TODO: get rid of all the debugging print statements
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
				fmt.Println(1)
				return nil
			}
			i++

			if i >= len(tokens) {
				fmt.Println(2)
				return nil
			}
			if _, ok := linesKeys[tokens[i]]; !ok {
				fmt.Println(3)
				return nil
			}
			linesOp := tokens[i]
			i++

			if i >= len(tokens) || tokens[i] != "[" {
				fmt.Println(4)
				return nil
			}
			i++

			if i >= len(tokens) {
				fmt.Println(5)
				return nil
			}
			linesThreshold, err := strconv.Atoi(tokens[i])
			if err != nil {
				fmt.Println(6)
				return nil
			}
			i++

			if i >= len(tokens) || tokens[i] != "]" {
				fmt.Println(7)
				return nil
			}
			i++

			if i >= len(tokens) || tokens[i] != "]" {
				fmt.Println(8)
				return nil
			}
			i++

			linesMap[linesOp] = linesThreshold
			qMap[qType] = linesMap
		case "text":
			textMap := make(map[string]interface{})

			if i >= len(tokens) || tokens[i] != "[" {
				fmt.Println(9)
				return nil
			}
			i++

			continueProcessing := true
			for continueProcessing {
				if i >= len(tokens) {
					fmt.Println(10)
					return nil
				}
				if _, ok := textKeys[tokens[i]]; !ok {
					fmt.Println(11)
					return nil
				}
				if _, ok := textMap[tokens[i]]; ok {
					fmt.Println(12)
					return nil
				}
				textKey := tokens[i]
				i++

				if i >= len(tokens) || tokens[i] != "[" {
					fmt.Println(13)
					return nil
				}
				i++

				if i >= len(tokens) {
					fmt.Println(14)
					return nil
				}
				switch textKey {
				case "val":
					if len(tokens[i]) < 2 || tokens[i][0] != '"' || tokens[i][len(tokens[i])-1] != '"' {
						fmt.Println(15)
						return nil
					}

					textMap[textKey] = tokens[i][1:len(tokens[i])-1]
				case "regex":
					isRegex, err := strconv.ParseBool(tokens[i])
					if err != nil {
						fmt.Println(16)
						return nil
					}

					textMap[textKey] = isRegex
				}
				i++

				if i >= len(tokens) || tokens[i] != "]" {
					fmt.Println(17)
					return nil
				}
				i++

				if i >= len(tokens) || !(tokens[i] == "]" || tokens[i] == ",") {
					fmt.Println(18)
					return nil
				} else if tokens[i] == "]" {
					continueProcessing = false
				}
				i++
			}

			if _, ok := textMap["val"]; !ok {
				fmt.Println(19)
				return nil
			}
			if _, ok := textMap["regex"]; !ok {
				fmt.Println(20)
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