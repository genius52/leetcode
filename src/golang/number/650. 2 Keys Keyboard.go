package number

import "math"

//Input: 3
//Output: 3
//Explanation:
//Intitally, we have one character 'A'.
//In step 1, we use Copy All operation.
//In step 2, we use Paste operation to get 'AA'.
//In step 3, we use Paste operation to get 'AAA'.
func dfs_minsteps(target int,cur_len int,total int,memo []int)int{
	if total > target{
		return math.MaxInt32
	}
	if total == target{
		return 0
	}
	if cur_len + total > target{
		return math.MaxInt32
	}
	paste := 1 + dfs_minsteps(target,cur_len,total + cur_len,memo)
	paste_copy := 2 + dfs_minsteps(target,total + cur_len,total + cur_len,memo)
	memo[total] = min_int(paste,paste_copy)
	return memo[total]
}

func MinSteps(n int) int {
	if n == 1{
		return 0
	}
	var memo []int = make([]int,n + 1)
	res := 1 + dfs_minsteps(n,1,1,memo)
	return res
}