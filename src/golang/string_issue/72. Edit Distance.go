package string_issue

import "math"

//Input: word1 = "horse", word2 = "ros"
//Output: 3
//Explanation:
//horse -> rorse (replace 'h' with 'r')
//rorse -> rose (remove 'r')
//rose -> ros (remove 'e')

func min_int_number(nums ...int)int{
	var min int = math.MaxInt32
	for _,n := range nums{
		if n < min{
			min = n
		}
	}
	return min
}

func dp_minDistance(word1 string,pos1 int,word2 string,pos2 int,record map[string]int)int{
	l1 := len(word1)
	l2 := len(word2)
	if pos1 == l1{
		return l2  - pos2
	}
	if pos2 == l2{
		return l1 - pos1
	}
	k := string(pos1) + "," + string(pos2)
	if _,ok := record[k];ok{
		return record[k]
	}
	//Insert a character
	//Delete a character
	//Replace a character
	equal_len := math.MaxInt32
	if word1[pos1] == word2[pos2]{
		equal_len = dp_minDistance(word1,pos1 + 1,word2,pos2 + 1,record)
	}
	delete_len := 1 + dp_minDistance(word1,pos1 + 1,word2,pos2,record)
	insert_len := 1 + dp_minDistance(word1,pos1,word2,pos2 + 1,record)
	replace_len := 1 + dp_minDistance(word1,pos1 + 1,word2,pos2 + 1,record)
	res := min_int_number(equal_len,delete_len,insert_len,replace_len)
	record[k] = res
	return res
}

func minDistance(word1 string, word2 string) int {
	var record map[string]int = make(map[string]int)
	return dp_minDistance(word1,0,word2,0,record)
}

func minDistance2(word1 string, word2 string) int {
	var len1 int = len(word1)
	var len2 int = len(word2)
	if len1 == 0{
		return len2
	}
	if len2 == 0{
		return len1
	}
	var dp [][]int = make([][]int,len1 + 1)//dp[i][j] = min steps of change word1[0:i] to word2[0:j]
	for i := 0;i <= len1;i++{
		dp[i] = make([]int,len2 + 1)
	}
	for i := 1;i <= len1;i++{
		dp[i][0] = i
	}
	for j := 1;j <= len2;j++{
		dp[0][j] = j
	}
	for i := 1;i <= len1;i++{
		for j := 1;j <= len2;j++{
			if word1[i - 1] == word2[j - 1]{
				dp[i][j] = dp[i - 1][j - 1]
			}else{
				dp[i][j] = min_int_number(dp[i][j - 1],dp[i - 1][j],dp[i - 1][j - 1]) + 1
			}
		}
	}
	return dp[len1][len2]
}
