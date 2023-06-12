package string_issue

import "strings"

func removeTrailingZeros(num string) string {
	return strings.TrimRight(num, "0")
}
