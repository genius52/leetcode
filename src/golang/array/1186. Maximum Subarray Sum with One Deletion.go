package array

//Input: arr = [1,-2,-2,3]
//Output: 3
func maximumSum(arr []int) int{
	var l int = len(arr)
	var prefix []int = make([]int,l)
	var suffix []int = make([]int,l)
	prefix[0] = arr[0]
	suffix[l - 1] = arr[l - 1]
	for i := 1;i < l;i++{
		prefix[i] = max_int_number(arr[i],arr[i] + prefix[i - 1])
	}
	for i := l - 2;i >= 0;i--{
		suffix[i] = max_int_number(arr[i],arr[i] + suffix[i + 1])
	}
	var res int = -2147483648
	for i := 0;i < l;i++{
		var cur int = 0
		if i == 0{
			cur = suffix[i]
		}else if i == (l - 1){
			cur = prefix[i]
		}else{
			cur = max_int_number(arr[i] + prefix[i - 1] + suffix[i + 1],prefix[i - 1] + suffix[i + 1],
				arr[i] + prefix[i - 1],arr[i] + suffix[i + 1],arr[i])
		}
		if cur > res{
			res = cur
		}
	}
	return res
}

func MaximumSum(arr []int) int {
	var l int = len(arr)
	if l == 1{
		return arr[0]
	}
	var left []int = make([]int,l + 1)
	var right []int = make([]int,l + 1)//right[i] = begin[i: l - 1]
	var max_num int = -2147483648
	if arr[0] > 0{
		left[0] = arr[0]
	}
	if arr[0] > max_num{
		max_num = arr[0]
	}
	if arr[l - 1] > 0{
		right[l - 1] = arr[l - 1]
	}
	for i := 1;i < l;i++{
		if arr[i] > max_num{
			max_num = arr[i]
		}
		if arr[i] >= 0{
			left[i] = left[i - 1] + arr[i]
		}else{
			if left[i - 1] + arr[i] > 0{
				left[i] = left[i - 1] + arr[i]
			}
		}
	}
	for i := l - 2;i >= 0;i--{
		if arr[i] >= 0{
			right[i] = right[i + 1] + arr[i]
		}else{
			if right[i + 1] + arr[i] > 0{
				right[i] = right[i + 1] + arr[i]
			}
		}
	}
	var res int = -2147483648
	for i := 0;i < l;i++{
		var total int = 0
		if i == 0{
			total = right[i + 1]
		}else if i == l - 1{
			total = left[i - 1]
		}else{
			if arr[i] >= 0{
				total = left[i - 1] + arr[i] + right[i + 1]
			}else{
				total = left[i - 1] + right[i + 1]
			}
		}
		if total > res{
			res = total
		}
	}
	if res == 0{
		return max_num
	}
	return res
}

//func dp_maximumSum(arr []int,l int,cur_pos int,can_delete bool,start bool,memo [][2]int)int{
//	if cur_pos >= l{
//		return 0
//	}
//	if start{
//		if arr[cur_pos] >= 0{
//			return arr[cur_pos] + dp_maximumSum(arr,l,cur_pos + 1,can_delete,start,memo)
//		}else{
//			if can_delete{
//				del := dp_maximumSum(arr,l,cur_pos + 1,false,start,memo)
//				not_del := arr[cur_pos] + dp_maximumSum(arr,l,cur_pos + 1,true,start,memo)
//				return max_int(del,not_del)
//			}else{
//				not_end := arr[cur_pos] + dp_maximumSum(arr,l,cur_pos + 1,false,start,memo)
//				if not_end > 0{
//					return not_end
//				}else{
//					return 0
//				}
//			}
//		}
//	}else{
//		do_start := arr[cur_pos] + dp_maximumSum(arr,l,cur_pos + 1,true,true,memo)
//		not_start := dp_maximumSum(arr,l,cur_pos + 1,true,false,memo)
//		return max_int(do_start,not_start)
//	}
//}
//
//func MaximumSum(arr []int) int {
//	var l int = len(arr)
//	var memo [][2]int = make([][2]int,l)
//	res := dp_maximumSum(arr,l,0,true,false,memo)
//	if res == 0{
//		var max int = -2147483648
//		for _,a := range arr{
//			if a > max{
//				max = a
//			}
//		}
//		return max
//	}else{
//		return res
//	}
//}