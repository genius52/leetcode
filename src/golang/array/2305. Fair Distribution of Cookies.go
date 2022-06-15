package array

//func check_distributeCookies(cookies []int,k int,target int)bool{
//	var people int = 1
//	var cur_cnt int = 0
//	for _,n := range cookies{
//		if cur_cnt + n > target{
//			cur_cnt = n
//			people++
//		}else{
//			cur_cnt += n
//		}
//	}
//	return people <= k
//}

func dfs_distributeCookies(cookies []int,l1 int,idx int,arr []int,l2 int)int{
	if idx >= l1{
		var res int = 0
		for _,n := range arr{
			if n > res{
				res = n
			}
		}
		return res
	}
	var res int = 2147483647
	for i := 0;i < l2;i++{
		arr[i] += cookies[idx]
		res = min_int(res,dfs_distributeCookies(cookies,l1,idx + 1,arr,l2))
		arr[i] -= cookies[idx]
	}
	return res
}

func DistributeCookies(cookies []int, k int) int {
	//sort.Ints(cookies)
	var arr []int = make([]int,k)
	return dfs_distributeCookies(cookies,len(cookies),0,arr,k)
}