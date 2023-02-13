package array

func check_minCapability(nums []int,l int,k int,target int)bool{
	var cnt int = 0
	for i := 0;i < l;{
		if nums[i] <= target{
			cnt++
			if cnt >= k{
				return true
			}
			i += 2
		}else{
			i++
		}
	}
	return false
}

func MinCapability(nums []int, k int) int {
	var l int = len(nums)
	var low int = 1
	var high int = 0
	for _,n := range nums{
		if n > high{
			high = n
		}
	}
	var res int = high
	for low < high{
		mid := (low + high)/2
		ret := check_minCapability(nums,l,k,mid)
		if ret {
			high = mid
			if mid < res{
				res = mid
			}
		}else{
			low = mid + 1
		}
	}
	return res
}

//DP solution
//func dp_minCapability(nums []int,l int,idx int,cur_val int,k int,memo map[string]int)int{
//	if idx >= l{
//		if k == 0{
//			return cur_val
//		}else{
//			return 1 << 31 - 1
//		}
//	}
//	if k == 0{
//		return cur_val
//	}
//	key := strconv.Itoa(idx) + "," + strconv.Itoa(k) + "," + strconv.Itoa(cur_val)
//	if val,ok := memo[key];ok{
//		return val
//	}
//
//	//stole current room
//	var val1 int = max_int(cur_val,nums[idx])
//	ret1 := dp_minCapability(nums,l,idx + 2,val1,k - 1,memo)
//	//skip current room
//	ret2 := dp_minCapability(nums,l,idx + 1,cur_val,k,memo)
//	memo[key] = min_int(ret1,ret2)
//	return memo[key]
//}
//
//func MinCapability(nums []int, k int) int {
//	var l int = len(nums)
//	var memo map[string]int = make(map[string]int)
//	return dp_minCapability(nums,l,0,0,k,memo)
//}