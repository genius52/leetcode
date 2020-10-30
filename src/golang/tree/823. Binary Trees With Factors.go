package tree

import "sort"

func NumFactoredBinaryTrees(A []int) int {
	sort.Ints(A)
	var l int = len(A)
	var exist map[int]bool = make(map[int]bool)
	var dp map[int]int = make(map[int]int)//record[i] 数字i为树根的二叉树个数
	for _,n := range A{
		exist[n] = true
		dp[n] = 1
	}
	var res int = 0
	for i := 0;i < l;i++{
		for j := 0;j < i;j++{
			if A[i] % A[j] == 0{
				n := A[i] / A[j]
				if _,ok := exist[n];ok{
					dp[A[i]] += dp[A[j]] * dp[n]
					dp[A[i]] = dp[A[i]] % 1000000007
				}
			}

		}
	}
	for _,cnt := range dp{
		res += cnt
		res = res % 1000000007
	}
	return res
}