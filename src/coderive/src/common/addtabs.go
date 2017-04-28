package common

import "fmt"

// AddTabs adds tabs for each line for pretty printing.
func AddTabs(lines []string) []string {
	for idx, line := range lines {
		lines[idx] = fmt.Sprintf("\t\t%s", line)
	}
	return lines
}
