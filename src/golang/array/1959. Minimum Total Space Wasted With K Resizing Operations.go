package array

func dp_minSpaceWastedKResizing(nums []int,pos int,k int,memo [][]int)int{
	if pos >= len(nums){
		return 0
	}
	if k < 0{
		return 1147483647
	}
	if memo[pos][k] != -1{
		return memo[pos][k]
	}
	var res int = 2147483647
	var sum int = 0
	var cur_size int = 0
	for i := pos;i < len(nums);i++{
		cur_size = max_int(cur_size,nums[i])
		sum += nums[i]
		wasted := (i - pos + 1) * cur_size - sum
		cur := wasted + dp_minSpaceWastedKResizing(nums,i + 1,k - 1,memo)
		if cur < res{
			res = cur
		}
	}
	memo[pos][k] = res
	return res
}

func MinSpaceWastedKResizing(nums []int, k int) int {
	var l int = len(nums)
	var memo [][]int = make([][]int,l + 1)
	for i := 0;i < l;i++{
		memo[i] = make([]int,k + 1)
		for j := 0;j <= l;j++{
			memo[i][j] = -1
		}
	}
	return dp_minSpaceWastedKResizing(nums,0,k,memo)
}