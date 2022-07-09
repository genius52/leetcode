package string_issue

func LongestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0{
		return ""
	}else if l == 1{
		return strs[0]
	}
	var res string = strs[0]
	for i := 1;i < l;i++{
		var j int = 0
		for j < len(strs[i]) && j < len(res){
			if strs[i][j] != res[j]{
				break
			}
			j++
		}
		res = res[0:j]
		if j == 0{
			return ""
		}
	}
	return res
}