package string_issue

import "strings"

//DP from top to bottom
func dp_isInterleave(s1 string,idx1 int,l1 int,s2 string,idx2 int,l2 int,s3 string,idx3 int,l3 int,visited [][]bool)bool{
	if idx3 == l3{
		return true
	}
	if idx1 == l1{
		ret := strings.Compare(s2[idx2:],s3[idx3:]) == 0
		visited[idx1][idx2] = true
		return ret
	}
	if idx2 == l2{
		ret := strings.Compare(s1[idx1:],s3[idx3:]) == 0
		visited[idx1][idx2] = true
		return ret
	}
	if visited[idx1][idx2]{
		return false
	}
	if s1[idx1] == s3[idx3] && s2[idx2] == s3[idx3]{
		ret1 := dp_isInterleave(s1,idx1 + 1,l1,s2,idx2,l2,s3,idx3 + 1,l3,visited)
		if ret1{
			return true
		}
		ret2 := dp_isInterleave(s1,idx1,l1,s2,idx2 + 1,l2,s3,idx3 + 1,l3,visited)
		if ret2{
			return true
		}
		visited[idx1][idx2] = true
		return false
	}else if s1[idx1] == s3[idx3]{
		ret := dp_isInterleave(s1,idx1 + 1,l1,s2,idx2,l2,s3,idx3 + 1,l3,visited)
		visited[idx1][idx2] = true
		return ret
	}else if s2[idx2] == s3[idx3]{
		ret := dp_isInterleave(s1,idx1,l1,s2,idx2 + 1,l2,s3,idx3 + 1,l3,visited)
		visited[idx1][idx2] = true
		return ret
	}else{
		visited[idx1][idx2] = true
		return false
	}
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	var l1 int = len(s1)
	var l2 int = len(s2)
	var l3 int = len(s3)
	if l1 + l2 != l3{
		return false
	}
	var visited [][]bool = make([][]bool,l1 + 1)
	for i := 0;i <= l1;i++{
		visited[i] = make([]bool,l2 + 1)
	}
	return dp_isInterleave(s1,0,l1,s2,0,l2,s3,0,l3,visited)
}