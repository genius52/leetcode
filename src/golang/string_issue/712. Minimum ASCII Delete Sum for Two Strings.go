package string_issue

import "math"

//Input: s1 = "delete", s2 = "leet"
//Output: 403
//Explanation: Deleting "dee" from "delete" to turn the string into "let",
//adds 100[d]+101[e]+101[e] to the sum.  Deleting "e" from "leet" adds 101[e] to the sum.
//At the end, both strings are equal to "let", and the answer is 100+101+101+101 = 403.
//If instead we turned both strings into "lee" or "eet", we would get answers of 433 or 417, which are higher.

func dp_MinimumDeleteSum(s1 string,pos1 int,s2 string,pos2 int,memo [][]int)int{
	var len1 int = len(s1)
	var len2 int = len(s2)
	if(pos1 >= len1){
		var rest int = 0
		for i := pos2;i < len2;i++{
			rest += int(s2[i])
		}
		return rest
	}
	if(pos2 >= len2){
		var rest int = 0
		for i := pos1;i < len1;i++{
			rest += int(s1[i])
		}
		return rest
	}
	if(memo[pos1][pos2] != -1){
		return memo[pos1][pos2]
	}
	var res int = math.MaxInt32
	if(s1[pos1] == s2[pos2]){
		res = int(math.Min(float64(res),float64(dp_MinimumDeleteSum(s1,pos1 + 1,s2,pos2 + 1,memo))))
		memo[pos1][pos2] = res
	}else{
		res = (int)(math.Min(float64(res),float64(int(s1[pos1]) + dp_MinimumDeleteSum(s1,pos1 + 1,s2,pos2,memo))))
		res = (int)(math.Min(float64(res),float64(int(s2[pos2]) + dp_MinimumDeleteSum(s1,pos1,s2,pos2 + 1,memo))))
		memo[pos1][pos2] = res
	}
	return res
}

func MinimumDeleteSum(s1 string, s2 string) int {
	var len1 int = len(s1)
	var len2 int = len(s2)
	var memo [][]int = make([][]int,len1)
	for i := 0;i < len1;i++{
		memo[i] = make([]int,len2)
		for j := 0;j < len2;j++{
			memo[i][j] = -1
		}
	}
	return dp_MinimumDeleteSum(s1,0,s2,0,memo)
}

func LCS_minimumDeleteSum(s1 string, s2 string) int {
	var len1 int = len(s1)
	var len2 int = len(s2)
	var graph [][]int = make([][]int,len1)
	for i := 0;i < len1;i++{
		graph[i] = make([]int,len2)
	}
	for i := 0;i < len1;i++{
		graph[i][0] = int(s1[i])
	}
	for j := 0;j < len2;j++{
		graph[0][j] = int(s2[j])
	}
	for i := 1;i < len1;i++{
		for j := 1;j < len2;j++{
			if(s1[i] == s2[j]){
				graph[i][j] = int(math.Max(float64(graph[i - 1][j]),float64(graph[i][j - 1]))) + 1
			}else{
				graph[i][j] = int(math.Max(float64(graph[i - 1][j]),float64(graph[i][j - 1])))
			}
		}
	}

}