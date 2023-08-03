package string_issue

import (
	"sort"
	"strings"
)

func combine_minimumString(s1 string, s2 string) string {
	if strings.Contains(s1, s2) {
		return s1
	}
	if strings.Contains(s2, s1) {
		return s2
	}
	var l1 int = len(s1)
	var l2 int = len(s2)
	var short_len int = min_int(l1, l2)
	for short_len >= 0 {
		ret := strings.HasSuffix(s1[l1-short_len:], s2[:short_len])
		if ret {
			return s1 + s2[short_len:]
		} else {
			short_len--
		}
	}
	return s1 + s2
}

func MinimumString(a string, b string, c string) string {
	//compare a+b+c,a+c+b,b+a+c,b+c+a,c+a+b,c+b+a
	s1 := combine_minimumString(combine_minimumString(a, b), c)
	s11 := combine_minimumString(c, combine_minimumString(a, b))
	s2 := combine_minimumString(combine_minimumString(a, c), b)
	s22 := combine_minimumString(b, combine_minimumString(a, c))
	s3 := combine_minimumString(combine_minimumString(b, a), c)
	s33 := combine_minimumString(c, combine_minimumString(b, a))
	s4 := combine_minimumString(combine_minimumString(b, c), a)
	s44 := combine_minimumString(a, combine_minimumString(b, c))
	s5 := combine_minimumString(combine_minimumString(c, a), b)
	s55 := combine_minimumString(b, combine_minimumString(c, a))
	s6 := combine_minimumString(combine_minimumString(c, b), a)
	s66 := combine_minimumString(a, combine_minimumString(c, b))
	var record []string = []string{s1, s2, s3, s4, s5, s6, s11, s22, s33, s44, s55, s66}
	//sort.Strings(record)
	sort.Slice(record, func(i, j int) bool {
		if len(record[i]) == len(record[j]) {
			return record[i] <= record[j]
		} else {
			return len(record[i]) < len(record[j])
		}
	})
	return record[0]
}
