package array

//DP from top to bottom with memory
func dp_stoneGameV(stoneValue []int,l int,left int,right int,prefix []int,memo [][]int)int{
	if left == right{
		return 0
	}
	if left + 1 == right{
		return min_int(stoneValue[left],stoneValue[right])
	}
	if memo[left][right] != 0{
		return memo[left][right]
	}
	var res int = 0
	for i := left;i < right;i++{
		left_sum := prefix[i + 1] - prefix[left]
		right_sum := prefix[right + 1] - prefix[i + 1]
		var cur int = 0
		if left_sum < right_sum{
			cur = left_sum + dp_stoneGameV(stoneValue,l,left,i,prefix,memo)
		}else if left_sum > right_sum{
			cur = right_sum + dp_stoneGameV(stoneValue,l,i + 1,right,prefix,memo)
		}else{
			cur = max_int(left_sum + dp_stoneGameV(stoneValue,l,left,i,prefix,memo),
				right_sum + dp_stoneGameV(stoneValue,l,i + 1,right,prefix,memo))
		}
		if cur > res{
			res = cur
		}
	}
	memo[left][right] = res
	return res
}

func StoneGameV(stoneValue []int) int {
	var l int = len(stoneValue)
	var prefix []int = make([]int,l + 1)
	for i := 0;i < l;i++{
		prefix[i + 1] = stoneValue[i] + prefix[i]
	}
	var memo [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		memo[i] = make([]int,l)
	}
	return dp_stoneGameV(stoneValue,l,0,l - 1,prefix,memo)
}