package string_issue

import "sort"

func MinDeletions(s string) int {
	var cnt [26]int
	for _,c := range s{
		cnt[c - 'a']++
	}
	var record []int
	for i := 0;i < 26;i++{
		if cnt[i] != 0{
			record = append(record,cnt[i])
		}
	}
	sort.Ints(record)
	var l int = len(record)
	var steps int = 0
	for i := l - 2;i >= 0;i--{
		if record[i] >= record[i + 1]{
			if record[i + 1] == 0{
				steps += record[i]
				record[i] = 0
			}else{
				steps += record[i] - record[i + 1] + 1
				record[i] = record[i + 1] - 1
			}
		}
	}
	return steps
}