package string_issue

import "sort"

func minAnagramLength(s string) int {
	var l int = len(s)
	for sub_len := 1; sub_len <= l/2; sub_len++ {
		if l%sub_len != 0 {
			continue
		}
		var sub []uint8 = make([]uint8, sub_len)
		for j := 0; j < sub_len; j++ {
			sub[j] = s[j]
		}
		sort.Slice(sub, func(i, j int) bool {
			return sub[i] < sub[j]
		})
		var fail bool = false
		for i := sub_len; i < l; i += sub_len {
			var cur []uint8 = make([]uint8, sub_len)
			for k := 0; k < sub_len; k++ {
				cur[k] = s[i+k]
			}
			sort.Slice(cur, func(i, j int) bool {
				return cur[i] < cur[j]
			})
			for idx := 0; idx < sub_len; idx++ {
				if cur[idx] != sub[idx] {
					fail = true
					break
				}
			}
			if fail {
				break
			}
		}
		if !fail {
			return sub_len
		}
	}
	return l
}
