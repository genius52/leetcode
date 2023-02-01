package array


func dfs_mincost(nums []int,l int,idx int,k int,memo []int)int{
	if idx == l{
		return 0
	}
	if memo[idx] != 1 << 31 - 1{
		return memo[idx]
	}
	var scores int = 0
	var record map[int]int = make(map[int]int)
	for i := idx;i < l;i++{
		record[nums[i]]++
		if record[nums[i]] == 2{
			scores += 2
		}else if record[nums[i]] > 2{
			scores++
		}
		memo[idx] = min_int(memo[idx],scores + k + dfs_mincost(nums,l,i + 1,k,memo))
	}
	return memo[idx]
}

func MinCost2547(nums []int, k int) int {
	var l int = len(nums)
	var memo []int = make([]int,l) //memo[i] 最小代价
	for i := 0;i < l;i++{
		memo[i] = 1 << 31 - 1
	}
	ret := dfs_mincost(nums,l,0,k,memo)
	return ret
}