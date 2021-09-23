package array

func recursive_maxJumps(arr []int,l int,pos int,pre_num int,d int,dp *[]int)int{
	if pos < 0 || pos >= l{
		return 0
	}
	if (*dp)[pos] != 0{
		return (*dp)[pos]
	}
	(*dp)[pos] = 1
	if pre_num != -1 && pre_num <= arr[pos]{
		return 0
	}
	for i := pos + 1;i < l && i <= (pos + d);i++{
		if arr[i] >= arr[pos]{
			break
		}
		(*dp)[pos] = max_int((*dp)[pos],1 + recursive_maxJumps(arr,l,i,arr[pos],d,dp))
	}
	for i := pos - 1;i >= 0 && i >= (pos - d);i--{
		if arr[i] >= arr[pos]{
			break
		}
		(*dp)[pos] = max_int((*dp)[pos],1 + recursive_maxJumps(arr,l,i,arr[pos],d,dp))
	}
	return (*dp)[pos]
}

func MaxJumps(arr []int, d int) int {
	var l int = len(arr)
	var dp []int = make([]int,l)//dp[i] i位置附近最大的长度
	var res int = 0
	for i := 0;i < l;i++{
		res = max_int(res,recursive_maxJumps(arr,l,i,-1,d,&dp))
	}
	return res
}