package array

func recursive_maxJumps(arr []int,l int,pos int,pre_num int,d int,dp *[]int,visited *[]bool)int{
	if pos < 0 || pos >= l{
		return 0
	}
	if (*visited)[pos]{
		return 0
	}
	(*visited)[pos] = true
	if (*dp)[pos] != 1{
		return (*dp)[pos]
	}
	if pre_num != -1 && pre_num <= arr[pos]{
		return 0
	}
	for i := pos + 1;i < l && i <= (pos + d);i++{
		if arr[i] >= arr[pos]{
			break
		}
		(*dp)[pos] = max_int((*dp)[pos],1 + recursive_maxJumps(arr,l,i,arr[pos],d,dp,visited))
	}
	for i := pos - 1;i >= 0 && i >= (pos - d);i--{
		if arr[i] >= arr[pos]{
			break
		}
		(*dp)[pos] = max_int((*dp)[pos],1 + recursive_maxJumps(arr,l,i,arr[pos],d,dp,visited))
	}
	(*visited)[pos] = false
	return (*dp)[pos]
}

func MaxJumps(arr []int, d int) int {
	var l int = len(arr)
	var dp []int = make([]int,l)//dp[i] i位置附近最大的长度
	for i := 0;i < l;i++{
		dp[i] = 1
	}
	var res int = 0
	for i := 0;i < l;i++{
		var visited []bool = make([]bool,l)
		res = max_int(res,recursive_maxJumps(arr,l,i,-1,d,&dp,&visited))
	}
	return res
}