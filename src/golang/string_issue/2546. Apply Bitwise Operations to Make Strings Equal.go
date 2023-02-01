package string_issue

import "strings"

//00 -> 00 不变
//01 -> 11 增加一个1
//11 -> 10 减少一个1
func makeStringsEqual(s string, target string) bool {
	if s == target{
		return true
	}
	if !strings.Contains(s,"1") || !strings.Contains(target,"1"){
		return false
	}
	return true
}