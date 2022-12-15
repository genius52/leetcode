package array

//func dfs_maxJump(stones []int,l int,idx int,forward bool,pre_max int,visited []bool)int{
//	if forward{
//
//	}else{
//
//	}
//}

func maxJump(stones []int) int {
	var l int = len(stones)
	var res int = stones[1] - stones[0]
	for i := 0;i + 2 < l;i += 2{
		dis := stones[i + 2] - stones[i]
		if dis > res{
			res = dis
		}
	}
	for i := 1;i + 2 < l;i += 2{
		dis := stones[i + 2] - stones[i]
		if dis > res{
			res = dis
		}
	}
	return res
}