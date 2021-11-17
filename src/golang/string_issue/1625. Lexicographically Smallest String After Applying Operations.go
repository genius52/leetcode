package string_issue

import (
	"strings"
)

func add_a(s string,l int,a int)string{
	tmp := []byte(s)
	for i := 1;i < l;i += 2{
		tmp[i] = '0' + (tmp[i] - '0' + uint8(a)) % 10
	}
	return string(tmp)
}

func dfs_findLexSmallestString(s string,l int,a int,b int,visited map[string]bool,res *string){
	if _,ok := visited[s];ok{
		return
	}
	visited[s] = true
	if strings.Compare(s,*res) < 0{
		*res = s
	}
	dfs_findLexSmallestString(s[b:] + s[:b],l,a,b,visited,res)
	dfs_findLexSmallestString(add_a(s,l,a),l,a,b,visited,res)
}

func FindLexSmallestString(s string, a int, b int) string {
	var l int = len(s)
	var visited map[string]bool = make(map[string]bool)
	var res string = s
	dfs_findLexSmallestString(s,l,a,b,visited,&res)
	return res
}

//func FindLexSmallestString(s string, a int, b int) string {
//
//	visited := make(map[string]bool)
//	res := s
//	add := func(s string) string {
//		temp := []byte(s)
//		for i := 1; i < len(s); i+=2 {
//			temp[i] = byte('0'+(int(temp[i]-'0')+a)%10)
//		}
//		return string(temp)
//	}
//	var dfs func(s string)
//	dfs = func(s string){
//		if visited[s] {
//			return
//		}
//		visited[s] = true
//		if s < res {
//			res = s
//		}
//		dfs(add(s))
//		dfs(s[b:] + s[:b])
//	}
//	dfs(s)
//
//	return res
//}