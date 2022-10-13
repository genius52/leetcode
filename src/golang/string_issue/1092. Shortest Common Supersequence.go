package string_issue

import "strings"

//Input: str1 = "abac", str2 = "cab"
//Output: "cabac"
func ShortestCommonSupersequence(str1 string, str2 string) string{
	var l1 int = len(str1)
	var l2 int = len(str2)
	var visited map[int]int = make(map[int]int)
	var pre string = str1 + str2
	var cur []uint8
	dp_shortestCommonSupersequence(str1,l1,0,str2,l2,0,&cur,visited,&pre)
	return pre
}

func dp_shortestCommonSupersequence(str1 string,l1 int,idx1 int,str2 string,l2 int,idx2 int,cur *[]uint8,visited map[int]int,pre *string){
	cur_len := len(*cur)
	min_len := max_int(l1 - idx1,l2 - idx2) + cur_len
	if min_len >= len(*pre){
		return
	}
	key := idx1 * 1000 + idx2

	if val,ok := visited[key];ok{
		if val <= cur_len{
			return
		}
	}
	visited[key] = cur_len
	if idx1 == l1{
		var ss strings.Builder
		for _,c := range *cur{
			ss.WriteByte(c)
		}
		for i := idx2;i < l2;i++{
			ss.WriteByte(str2[i])
		}
		cur := ss.String()
		if len(cur) < len(*pre){
			*pre = cur
		}
		return
	}
	if idx2 == l2{
		var ss strings.Builder
		for _,c := range *cur{
			ss.WriteByte(c)
		}
		for i := idx1;i < l1;i++{
			ss.WriteByte(str1[i])
		}
		cur := ss.String()
		if len(cur) < len(*pre){
			*pre = cur
		}
		return
	}
	if str1[idx1] == str2[idx2]{
		*cur = append(*cur,str1[idx1])
		dp_shortestCommonSupersequence(str1,l1,idx1 + 1,str2,l2,idx2 + 1,cur,visited,pre)
		*cur = (*cur)[:len(*cur) - 1]
	}else{
		*cur = append(*cur,str1[idx1])
		dp_shortestCommonSupersequence(str1,l1,idx1 + 1,str2,l2,idx2,cur,visited,pre)
		*cur = (*cur)[:len(*cur) - 1]
		*cur = append(*cur,str2[idx2])
		dp_shortestCommonSupersequence(str1,l1,idx1,str2,l2,idx2 + 1,cur,visited,pre)
		*cur = (*cur)[:len(*cur) - 1]
	}
}

func ShortestCommonSupersequence2(str1 string, str2 string) string {
	var l1 int = len(str1)
	var l2 int = len(str2)
	var dp [][]string = make([][]string,l1 + 1)//dp[i + 1][j + 1]; max length common string from str1[0:i],str2[0:j]
	for i := 0;i <= l1;i++{
		dp[i] = make([]string,l2 + 1)
	}
	//calculate lcs
	for i := 0;i < l1;i++{
		for j := 0;j < l2;j++{
			if str1[i] == str2[j]{
				dp[i + 1][j + 1] = dp[i][j] + string(str1[i])
			}else{
				//dp[i + 1][j + 1] = dp[i + 1][j].size() > dp[i][j + 1].size() ?  dp[i + 1][j] : dp[i][j + 1];
				if len(dp[i + 1][j]) > len(dp[i][j + 1]){
					dp[i + 1][j + 1] = dp[i + 1][j]
				}else{
					dp[i + 1][j + 1] = dp[i][j + 1]
				}
			}
		}
	}
	var lcs string = dp[l1][l2]
	var ss strings.Builder
	var idx1 int = 0
	var idx2 int = 0
	for i := 0;i < len(lcs);i++{
		for idx1 < l1 && str1[idx1] != lcs[i]{
			ss.WriteByte(str1[idx1])
			idx1++
		}
		for idx2 < l2 && str2[idx2] != lcs[i]{
			ss.WriteByte(str2[idx2])
			idx2++
		}
		ss.WriteByte(lcs[i])
		idx1++
		idx2++
	}
	for idx1 < l1{
		ss.WriteByte(str1[idx1])
		idx1++
	}
	for idx2 < l2{
		ss.WriteByte(str2[idx2])
		idx2++
	}
	return ss.String()
}