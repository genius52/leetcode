package string_issue

func partitionString(s string) int {
	var l int = len(s)
	var res int = 1
	var visited [26]bool
	for i := 0;i < l;i++{
		if visited[s[i] - 'a']{
			res++
			for j := 0;j < 26;j++{
				visited[j] = false
			}
			visited[s[i] - 'a'] = true
		}else{
			visited[s[i] - 'a'] = true
		}
	}
	return res
}