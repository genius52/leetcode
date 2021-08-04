package array

import "sort"

type Start_end [][]int

func (s Start_end) Less(i, j int) bool {
	if s[i][0] < s[j][0]{
		return true
	}else if s[i][0] == s[j][0]{
		return s[i][1] > s[j][1]
	}
	return false
}

func (s Start_end) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Start_end) Len() int {
	return len(s)
}

//greedy
func VideoStitching(clips [][]int, T int) int {
	sort.Sort(Start_end(clips))
	var cnt int = 0
	var pre_end int = 0
	var l int = len(clips)
	for i := 0;i < l;{
		var furtherest int = 0
		j := i
		for ;j < l;j++{
			if clips[j][0] > pre_end {
				break
			}
			if clips[j][1] > furtherest{
				furtherest = clips[j][1]
			}
		}
		cnt++
		if furtherest == 0{
			return -1
		}
		if furtherest >= T{
			return cnt
		}
		pre_end = furtherest
		i = j
	}
	return -1
}

//DP solution
//func dfs_videoStitching(clips [][]int,l int,pos int,pre_end int,T int,memo map[int]int)int{
//	if pre_end >= T{
//		return 0
//	}
//	if _,ok := memo[pre_end];ok{
//		return memo[pre_end]
//	}
//	var res int = 2147483647
//	for i := pos;i < l;i++{
//		if clips[i][1] <= pre_end{
//			continue
//		}
//		if clips[i][0] <= pre_end{
//			if clips[i][1] >= T{
//				res = 1
//				break
//			}
//			cnt := dfs_videoStitching(clips,l,i + 1,clips[i][1],T,memo)
//			if cnt != -1{
//				if (cnt + 1) < res{
//					res = cnt + 1
//				}
//			}
//		}
//	}
//
//	if res == 2147483647{
//		memo[pre_end] = -1
//		return -1
//	}
//	memo[pre_end] = res
//	return res
//}
//
//func VideoStitching(clips [][]int, T int) int {
//	sort.Sort(Start_end(clips))
//	var memo map[int]int = make(map[int]int)
//	return dfs_videoStitching(clips,len(clips),0,0,T,memo)
//}