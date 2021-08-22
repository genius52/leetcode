package array

import "sort"

func abs_diff(a int,b int)int{
	if a > b{
		return a - b
	}else{
		return b - a
	}
}

func dp_minimizeTheDifference(mat *[][]int,rows int,columns int,r int,choose *[]int,cur_sum int,target int,memo *map[int]map[int]int)(int,int){
	if cur_sum == target{
		return 0,cur_sum
	}
	if r >= rows{
		return abs_diff(cur_sum,target),cur_sum
	}
	keys := r * 100 + (*choose)[r]
	if _,ok1 := (*memo)[keys];ok1{
		if _,ok2 := (*memo)[keys][cur_sum];ok2{
			return (*memo)[keys][cur_sum],cur_sum
		}
	}
	//skip current row
	res1,sum1 := dp_minimizeTheDifference(mat,rows,columns,r + 1,choose,cur_sum,target,memo)
	//move current row to right if available

	var res2 int = 2147483647
	var sum2 int = 0
	if ((*choose)[r] + 1) < columns{
		for i := 1;i < columns && ((*choose)[r] + i) < columns;i++{
			(*choose)[r] += i
			cur,s2 := dp_minimizeTheDifference(mat,rows,columns,r + 1,choose,cur_sum + (*mat)[r][(*choose)[r]] - (*mat)[r][(*choose)[r] - i],target,memo)
			(*choose)[r] -= i
			if cur < res2{
				res2 = cur
				sum2 = s2
			}else{
				if s2 > target{
					break
				}
			}
		}
	}
	if _,ok := (*memo)[keys];!ok{
		(*memo)[keys] = make(map[int]int)
	}
	if res1 < res2{
		(*memo)[keys][cur_sum] = res1
		return res1,sum1
	}else{
		(*memo)[keys][cur_sum] = res2
		return res2,sum2
	}
}

func minimizeTheDifference(mat [][]int, target int) int {
	var rows int = len(mat)
	var columns int = len(mat[0])
	for i := 0;i < rows;i++{
		sort.Ints(mat[i])
	}
	low := 0
	high := 0
	for i := 0;i < rows;i++{
		low += mat[i][0]
		high += mat[i][columns - 1]
	}
	if low >= target{
		return low - target
	}
	if high <= target{
		return target - high
	}
	var choose []int = make([]int,rows)
	var memo map[int]map[int]int = make(map[int]map[int]int)//rows * 100 + columns,sum
	res,_ := dp_minimizeTheDifference(&mat,rows,columns,0,&choose,low,target,&memo)
	return res
}