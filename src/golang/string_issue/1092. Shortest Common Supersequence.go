package string_issue

//Input: str1 = "abac", str2 = "cab"
//Output: "cabac"
func dp_shortestCommonSupersequence(str1 string,l1 int,pos1 int,str2 string,l2 int,pos2 int,memo *map[int]string)string{
	if pos1 >= l1 && pos2 >= l2{
		return ""
	}
	if pos1 >= l1{
		return str2[pos2:]
	}
	if pos2 >= l2{
		return str1[pos1:]
	}
	if _,ok := (*memo)[pos1 * 1000 + pos2];ok{
		return (*memo)[pos1 * 1000 + pos2]
	}
	var res string
	if str1[pos1] == str2[pos2]{
		res = string(str1[pos1]) + dp_shortestCommonSupersequence(str1,l1,pos1 + 1,str2,l2,pos2 + 1,memo)
	}else{
		res1 := string(str1[pos1]) + dp_shortestCommonSupersequence(str1,l1,pos1 + 1,str2,l2,pos2,memo)
		res2 := string(str2[pos2]) + dp_shortestCommonSupersequence(str1,l1,pos1,str2,l2,pos2 + 1,memo)
		if len(res1) < len(res2){
			res = res1
		}else{
			res = res2
		}
	}
	(*memo)[pos1 * 1000 + pos2] = res
	return res
}

func ShortestCommonSupersequence(str1 string, str2 string) string {
	var l1 int = len(str1)
	var l2 int = len(str2)
	var memo map[int]string = make(map[int]string)
	return dp_shortestCommonSupersequence(str1,l1,0,str2,l2,0,&memo)
}


func ShortestCommonSupersequence2(str1 string, str2 string) string {
	var l1 int = len(str1)
	var l2 int = len(str2)
	var dp [][]string = make([][]string,l1 + 1)//dp[i][j]; min length from str1[0:i],str2[0:j]
	for i := 0;i <= l1;i++{
		dp[i] = make([]string,l2 + 1)
	}
	for i := 1;i <= l1;i++{
		dp[i][0] = str1[0:i]
	}
	for i := 1;i <= l2;i++{
		dp[0][i] = str2[0:i]
	}
	//calculate lcs
	for i := 0;i < l1;i++{
		for j := 0;j < l2;j++{
			if str1[i] == str2[j]{
				dp[i + 1][j + 1] = dp[i][j] + string(str1[i])
			}else{
				if len(dp[i + 1][j]) < len(dp[i][j + 1]){
					dp[i + 1][j + 1] = dp[i + 1][j] + string(str1[i])
				}else{
					dp[i + 1][j + 1] = dp[i][j + 1] + string(str2[j])
				}
			}
		}
	}
	return dp[l1][l2]
}