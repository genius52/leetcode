package string_issue

import (
	"sort"
	"strings"
)

func frequencySort(s string) string {
	var char_cnt map[int]int = make(map[int]int)
	for _,c := range s{
		char_cnt[int(c)]++
	}
	var cnt_char [][2]int
	for k,v := range char_cnt{
		cnt_char = append(cnt_char,[2]int{v,k})
	}
	sort.Slice(cnt_char, func(i, j int) bool {
		return cnt_char[i][0] > cnt_char[j][0]
	})
	var ss strings.Builder
	for i := 0;i < len(cnt_char);i++{
		c := cnt_char[i][1]
		for j := 0;j < cnt_char[i][0];j++{
			ss.WriteRune(int32(c))
		}
	}
	return ss.String()
}