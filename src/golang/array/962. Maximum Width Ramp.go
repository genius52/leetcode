package array

import (
	"container/list"
	"math"
)

//Input: [6,0,8,2,1,5]
//Output: 4
//Explanation:
//The maximum width ramp is achieved at (i, j) = (1, 5): A[1] = 0 and A[5] = 5.

func MaxWidthRamp2(A []int) int{
	var l int = len(A)
	if l <= 1{
		return 0
	}
	var res int = 0
	var q list.List
	for i := 0;i < l;i++{
		if q.Len() == 0{
			q.PushBack(i)
		}else{
			top := q.Back().Value.(int)
			if A[i] < A[top]{
				q.PushBack(i)
			}else{
				visit := q.Front()
				for A[visit.Value.(int)] > A[i]{
					visit = visit.Next()
				}
				dis := i - visit.Value.(int)
				if dis > res{
					res = dis
				}
			}
		}
	}
	return res
}

func MaxWidthRamp(A []int) int{
	var l int = len(A)
	if l <= 1{
		return 0
	}
	var res int = 0
	for i := 0;i < l;i++{
		if res > l - 1 - i{
			break
		}
		for j := l - 1;j >= i;j--{
			if A[i] <= A[j]{
				res = int(math.Max(float64(res),float64(j - i)))
				break
			}
			if res > j - i{
				break
			}
		}
	}
	return res
}

//func dp_maxWidthRamp(A []int,left int,right int,memo map[string]int)int{
//	if left == right{
//		return 0
//	}
//	if A[left] <= A[right]{
//		return right - left
//	}
//	k := strconv.Itoa(left) + "," + strconv.Itoa(right)
//	if val,ok := memo[k];ok{
//		return val
//	}
//
//	move_left := dp_maxWidthRamp(A,left + 1,right,memo)
//	move_right := dp_maxWidthRamp(A,left,right - 1,memo)
//	memo[k] = int(math.Max(float64(move_left),float64(move_right)))
//	return memo[k]
//}
//
//func MaxWidthRamp(A []int) int{
//	var l int = len(A)
//	if l <= 1{
//		return 0
//	}
//	var memo map[string]int = make(map[string]int)
//	return dp_maxWidthRamp(A,0,l - 1,memo)
//}

//func dfs_maxWidthRamp(left_min []int,right_max []int,left int,right int)int{
//	if left == right{
//		return 0
//	}
//	if
//}

//func maxWidthRamp(A []int) int {
//	var l int = len(A)
//	if l <= 1{
//		return 0
//	}
//	var left_min []int = make([]int,l)
//	left_min[0] = A[0]
//	var right_max []int = make([]int,l)
//	right_max[l - 1] = A[l - 1]
//	for i := 1;i < l;i++{
//		if A[i] < left_min[i - 1]{
//			left_min[i] = A[i]
//		}else{
//			left_min[i] = left_min[i - 1]
//		}
//		if A[l - 1 - i] > A[l - i]{
//			right_max[l - 1 - i] = A[l - 1 - i]
//		}else{
//			right_max[l - 1 - i] = right_max[l - i]
//		}
//	}
//	return dfs_maxWidthRamp(left_min,right_max,0,l - 1)
//}