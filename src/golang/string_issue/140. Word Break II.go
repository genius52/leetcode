package string_issue

import "strings"

//140
//Input:
//s = "pineapplepenapple"
//wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
//Output:
//[
//  "pine apple pen apple",
//  "pineapple pen apple",
//  "pine applepen apple"
//]
func dp_wordBreak(s string,wordDict []string,sub string,record *[]string) {
	if len(s) == 0{
		*record = append(*record,sub)
		return
	}
	l := len(wordDict)
	for i := 0;i < l;i++{
		if !strings.HasPrefix(s,wordDict[i]){
			continue
		}
		var newsub string = sub + " "+ wordDict[i]
		dp_wordBreak(s[len(wordDict[i]):],wordDict,newsub,record)
	}
}

func WordBreak2(s string, wordDict []string) []string {
	l := len(wordDict)
	var res []string
	if l == 0{
		return res
	}
	if len(s) == 0{
		return res
	}
	for i := 0;i < l;i++{
		if len(wordDict[i]) == 0{
			continue
		}
		if !strings.HasPrefix(s,wordDict[i]){
			continue
		}
		var sub string = wordDict[i]
		dp_wordBreak(s[len(wordDict[i]):],wordDict,sub,&res)
	}
	return res
}