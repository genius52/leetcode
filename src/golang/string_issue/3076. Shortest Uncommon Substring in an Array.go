package string_issue

import "strings"

func ShortestSubstrings(arr []string) []string {
	var l int = len(arr)
	var visited map[string]int = make(map[string]int)
	for i := 0; i < l; i++ {
		visited[arr[i]]++
	}
	var res []string = make([]string, l)
	for i := 0; i < l; i++ {
		var str_len int = len(arr[i])
		visited[arr[i]]--
		if visited[arr[i]] == 0 {
			delete(visited, arr[i])
		}
		for sub_len := 1; sub_len <= str_len; sub_len++ {
			for start := 0; start+sub_len <= str_len; start++ {
				var sub string = arr[i][start : start+sub_len]
				var find bool = false
				for s, _ := range visited {
					if strings.Contains(s, sub) {
						find = true
						break
					}
				}
				if find {
					continue
				}
				if len(res[i]) == 0 {
					res[i] = sub
				} else {
					if strings.Compare(sub, res[i]) <= 0 {
						res[i] = sub
					}
				}
			}
			if len(res[i]) > 0 {
				break
			}
		}
		visited[arr[i]]++
	}
	return res
}
