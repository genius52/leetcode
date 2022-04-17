package string_issue

func dfs_generateParenthesis(n int,left int,right int,s string,res *[]string){
	if len(s) == n * 2{
		if left == right{
			*res = append(*res,s)
			return
		}
	}
	if left == n{
		dfs_generateParenthesis(n,left,right + 1,s + ")",res)
	}else{
		if left == right{
			dfs_generateParenthesis(n,left + 1,right,s + "(",res)
		}else{
			dfs_generateParenthesis(n,left + 1,right,s + "(",res)
			dfs_generateParenthesis(n,left,right + 1,s + ")",res)
		}
	}
}

func generateParenthesis(n int) []string {
	var res []string
	dfs_generateParenthesis(n,0,0,"",&res)
	return res
}