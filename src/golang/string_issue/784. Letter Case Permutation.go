package string_issue

//Input: s = "a1b2"
//Output: ["a1b2","a1B2","A1b2","A1B2"]
func dfs_letterCasePermutation(s string,l int,pos int,res *[]string){
	if pos >= l{
		*res = append(*res,s)
		return
	}
	if s[pos] >= '0' && s[pos] <= '9'{
		dfs_letterCasePermutation(s,l,pos + 1,res)
		return
	}
	if s[pos] >= 'a' && s[pos] <= 'z'{
		data := []byte(s)
		data[pos] = s[pos] - 32
		dfs_letterCasePermutation(s,l,pos + 1,res)
		var s2 string = string(data)
		dfs_letterCasePermutation(s2,l,pos + 1,res)
	}else{
		data := []byte(s)
		data[pos] = s[pos] + 32
		dfs_letterCasePermutation(s,l,pos + 1,res)
		var s2 string = string(data)
		dfs_letterCasePermutation(s2,l,pos + 1,res)
	}
}

func letterCasePermutation(s string) []string {
	var res []string
	var l int = len(s)
	dfs_letterCasePermutation(s,l,0,&res)
	return res
}