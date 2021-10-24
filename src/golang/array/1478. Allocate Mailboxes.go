package array

import "sort"

func dp_minDistance(cost [][]int,l int,pos int,k int,memo [][]int)int{
	if k == l - pos{
		return 0
	}
	//1.邮筒比房子多, 2.邮筒没了
	if k > l - pos || k == 0{ //this condition is not satisfied
		return 2147483647
	}
	if memo[pos][k] != 0{
		return memo[pos][k]
	}
	var res int = 2147483647
	for i := pos;i < l;i++{
		//放一个邮筒在位置i
		cur := cost[pos][i] + dp_minDistance(cost,l,i + 1,k - 1,memo)
		if cur < res{
			res = cur
		}
	}
	memo[pos][k] = res
	return res
}

func MinDistance(houses []int, k int) int {
	var l int = len(houses)
	sort.Ints(houses)
	var cost [][]int = make([][]int,l)//cost[i][j]: 在i和j之间放入一个邮筒，所产生的距离之和
	for i := 0;i < l;i++{
		cost[i] = make([]int,l)
	}
	for i := 0;i < l;i++{
		for j := 0;j < l;j++{
			mid := (i + j)/2
			for k := i;k <= j;k++{
				cost[i][j] += abs_int(houses[mid] - houses[k])
			}
		}
	}
	var memo [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		memo[i] = make([]int,k + 1)
	}
	return dp_minDistance(cost,l,0,k,memo)
}