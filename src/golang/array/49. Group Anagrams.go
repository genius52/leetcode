package array

import "sort"

//Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
//Output:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	var m map[string][]string
	for i := 0;i < len(strs);i++{
		data := []rune(strs[i])
		sort.Sort(sortRunes(data))
		var s string = string(data[:])
		m[s] = append(m[s],strs[i])
	}
	var res [][]string
	for _,val := range m{
		res = append(res,val)
	}
	return res
}