package string_issue

import (
	"sort"
	"strings"
)

func arrangeWords(text string) string {
	var record []string = strings.Split(strings.ToLower(text), " ")
	sort.SliceStable(record, func(i, j int) bool {
		var l1 int = len(record[i])
		var l2 int = len(record[j])
		return l1 < l2
	})
	record[0] = string(record[0][0] - 32) + record[0][1:]
	return strings.Join(record," ")
}