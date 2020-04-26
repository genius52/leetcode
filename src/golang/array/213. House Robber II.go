package array

//Input: [1,2,3,1]
//Output: 4
//Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
//             Total amount you can rob = 1 + 3 = 4.
func dp_rob(nums []int,l int, pos int,select_first bool,record_selectfirst map[int]int,record_notselectfirst map[int]int)int{
	if pos >= l{
		return 0
	}

	if select_first{
		if _,ok := record_selectfirst[pos];ok{
			return record_selectfirst[pos]
		}
	}else{
		if _,ok := record_notselectfirst[pos];ok{
			return record_notselectfirst[pos]
		}
	}

	if pos == l - 1{
		if !select_first{
			return nums[pos]
		}else{
			return 0
		}
	}

	res := max_int(nums[pos] + dp_rob(nums,l,pos + 2,select_first,record_selectfirst,record_notselectfirst),
		dp_rob(nums,l,pos + 1,select_first,record_selectfirst,record_notselectfirst))
	if select_first{
		record_selectfirst[pos] = res
	}else{
		record_notselectfirst[pos] = res
	}
	return res
}

func Rob(nums []int) int {
	l := len(nums)
	if l == 0{
		return 0
	}
	var record_selectfirst map[int]int = make(map[int]int)
	var record_notselectfirst map[int]int = make(map[int]int)
	return max_int(nums[0] + dp_rob(nums,l,2,true,record_selectfirst,record_notselectfirst),
		dp_rob(nums,l,1,false,record_selectfirst,record_notselectfirst))
}

//dp no recursive
func dp_rob2(nums []int,low int,high int)int{
	var rob_last_last int = 0
	var rob_last int = nums[low]
	for i := low + 1;i <= high;i++{
		tmp := max_int(rob_last_last + nums[i],rob_last)
		rob_last_last = rob_last
		rob_last = tmp
	}
	return rob_last
}

func rob2(nums []int) int {
	l := len(nums)
	if l == 0{
		return 0
	}
	if l == 1{
		return nums[0]
	}
	return max_int(dp_rob2(nums,0,l - 2),dp_rob2(nums,1,l - 1))
}