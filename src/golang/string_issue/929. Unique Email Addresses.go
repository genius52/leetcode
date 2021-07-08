package string_issue

import "strings"

func numUniqueEmails(emails []string) int {
	var record map[string]bool = make(map[string]bool)
	var l int = len(emails)
	for i := 0;i < l;i++{
		data := strings.Split(emails[i],"@")
		prefix := data[0]
		pos := strings.IndexByte(prefix,'+')
		if pos != -1{
			prefix = prefix[:pos]
		}
		var cur string
		for _,s := range prefix{
			if s != '.'{
				cur += string(s)
			}
		}
		cur += "@" + data[1]
		record[cur] = true
	}
	return len(record)
}