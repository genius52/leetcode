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

//func dp_wordBreak(s string,wordDict []string,sub string,record *[]string) {
//	if len(s) == 0{
//		*record = append(*record,sub)
//		return
//	}
//	l := len(wordDict)
//	for i := 0;i < l;i++{
//		if !strings.HasPrefix(s,wordDict[i]){
//			continue
//		}
//		var newsub string = sub + " "+ wordDict[i]
//		dp_wordBreak(s[len(wordDict[i]):],wordDict,newsub,record)
//	}
//}
//
//func WordBreak2(s string, wordDict []string) []string {
//	l := len(wordDict)
//	var res []string
//	if l == 0{
//		return res
//	}
//	if len(s) == 0{
//		return res
//	}
//	for i := 0;i < l;i++{
//		if len(wordDict[i]) == 0{
//			continue
//		}
//		if !strings.HasPrefix(s,wordDict[i]){
//			continue
//		}
//		var sub string = wordDict[i]
//		dp_wordBreak(s[len(wordDict[i]):],wordDict,sub,&res)
//	}
//	return res
//}

func dp_wordBreak(s string,l int,idx int,wordDict []string,choose *[]int, res *[]string){
	if idx == l{
		var s2 string
		for _,c := range *choose{
			if len(s2) > 0{
				s2 += " "
			}
			s2 += wordDict[c]
		}
		*res = append(*res,s2)
		return
	}
	for i,word := range wordDict{
		if strings.HasPrefix(s[idx:],word){
			*choose = append(*choose,i)
			dp_wordBreak(s,l,idx + len(word),wordDict,choose,res)
			*choose = (*choose)[:len(*choose) - 1]
		}
	}
}

//To be optimized by memory

func WordBreak140(s string, wordDict []string) []string {
	var l int = len(s)
	var dict map[string]bool = make(map[string]bool)
	for _,w := range wordDict{
		dict[w] = true
	}
	var res []string
	var cur []int
	dp_wordBreak(s,l,0,wordDict,&cur,&res)
	return res
}