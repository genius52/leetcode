package string_issue

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0{
		return ""
	}else if l == 1{
		return strs[0]
	}
	var res string = strs[0]
	for i := 1;i < l;i++{
		for j:= 0;j < len(strs[i]) && j < len(res);j++{
			if strs[i][j] != res[j]{
				if j == 0{
					return ""
				}
				res = res[0:j]
				break
			}
		}
	}
	return res
}