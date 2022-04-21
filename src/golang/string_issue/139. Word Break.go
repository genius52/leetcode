package string_issue

import "strings"

//Input: s = "applepenapple", wordDict = ["apple", "pen"]
//Output: true
//Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
//Note that you are allowed to reuse a dictionary word.
//Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//Output: false
func dfs_wordBreak(s string,dict map[string]bool,cur_pos int,memo map[int]bool)bool{
	if cur_pos >= len(s){
		return true
	}
	if _,ok := memo[cur_pos];ok{
		return false
	}
	sub := s[cur_pos:]
	for w,_ := range dict{
		if strings.HasPrefix(sub,w){
			if dfs_wordBreak(s,dict,cur_pos + len(w),memo){
				return true
			}
		}
	}
	memo[cur_pos] = false
	return false
}

func wordBreak(s string, wordDict []string) bool {
	var record map[int]bool = make(map[int]bool)
	var dict map[string]bool = make(map[string]bool)
	for _,w := range wordDict{
		dict[w] = true
	}
	return dfs_wordBreak(s,dict,0,record)
}