package string_issue

//Input: s = "babgbag", t = "bag"
//Output: 5
//Explanation:
//As shown below, there are 5 ways you can generate "bag" from S.
//babgbag
//babgbag
//babgbag
//babgbag
//babgbag
func dp_numDistinct(s string,pos1 int,t string,pos2 int,memo [][]int)int{
	var len_t int = len(t)
	if pos2 >= len_t{
		return 1
	}
	var len_s int = len(s)
	if pos1 >= len_s{
		return 0
	}
	if (len_s - pos1) < (len_t - pos2){
		return 0
	}
	if memo[pos1][pos2] != -1{
		return memo[pos1][pos2]
	}
	var res int = 0
	if s[pos1] == t[pos2]{
		res += dp_numDistinct(s,pos1 + 1,t,pos2 + 1,memo)
		res += dp_numDistinct(s,pos1 + 1,t,pos2,memo)
	}else{
		res += dp_numDistinct(s,pos1 + 1,t,pos2,memo)
	}
	if res >= 0{
		memo[pos1][pos2] = res
	}
	return res
}

func NumDistinct(s string, t string) int {
	var l1 int = len(s)
	var l2 int = len(t)
	if l1 < l2{
		return 0
	}
	if l2 == 0{
		return 1
	}
	var memo [][]int = make([][]int,l1)
	for i := 0;i < l1;i++{
		memo[i] = make([]int,l2)
		for j := 0;j < l2;j++{
			memo[i][j] = -1
		}
	}
	dp_numDistinct(s,0,t,0,memo)
	return memo[0][0]
}