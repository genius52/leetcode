package array

func GetMaxLen(nums []int) int{
	var l int = len(nums)
	var dp_neg []int = make([]int,l)//dp_neg[i]： 当前乘积是负数的最大长度
	var dp_pos []int = make([]int,l)
	var res int = 0
	for i := 0;i < l;i++{
		if nums[i] == 0{
			dp_neg[i] = 0
			dp_pos[i] = 0
		}else if nums[i] < 0{
			if i == 0{
				dp_neg[i] = 1
			}else{
				if dp_pos[i - 1] > 0{
					dp_neg[i] = 1 + dp_pos[i - 1]
				}else{
					dp_neg[i] = 1
				}
				if dp_neg[i - 1] > 0{
					dp_pos[i] = 1 + dp_neg[i - 1]
				}
			}
		}else{
			if i == 0{
				dp_pos[i] = 1
			}else{
				if dp_neg[i - 1] > 0{
					dp_neg[i] = 1 + dp_neg[i - 1]
				}
				dp_pos[i] = 1 + dp_pos[i - 1]
			}
		}
		res = max_int(res,dp_pos[i])
	}
	return res
}

func GetMaxLen2(nums []int) int{
	var l int = len(nums)
	var pre_neg_len int = 0//上一个位置乘积为负数的最大长度
	var pre_pos_len int = 0//上一个位置乘积为正数的最大长度
	if nums[0] > 0{
		pre_pos_len = 1
	}else if nums[0] < 0{
		pre_neg_len = 1
	}
	var res int = 0
	for i := 1;i < l;i++{
		if nums[i] == 0{
			pre_neg_len = 0
			pre_pos_len = 0
		}else if nums[i] > 0{
			if pre_neg_len > 0{
				pre_neg_len++
			}
			pre_pos_len++
		}else if nums[i] < 0{
			tmp := pre_neg_len
			pre_neg_len = 1 + pre_pos_len
			if tmp > 0{
				pre_pos_len = tmp + 1
			}else{
				pre_pos_len = 0
			}
		}
		if pre_pos_len > res{
			res = pre_pos_len
		}
	}
	return res
}

//func GetMaxLen(nums []int) int {
//	nums = append(nums,0)
//	var l int = len(nums)
//	var begin int = 0
//	var end int = 0
//	var max_length int = 0
//	var neg_cnt int = 0
//	for end < l{
//		if nums[end] < 0{
//			neg_cnt++
//		}else if nums[end] == 0{
//			if neg_cnt % 2 == 0{
//				max_length = int(math.Max(float64(max_length),float64(end - begin)))
//			}else{
//				for begin < end{
//					if nums[begin] < 0{
//						max_length = int(math.Max(float64(max_length),float64(end - begin - 1)))
//						break
//					}
//					begin++
//				}
//			}
//			neg_cnt = 0
//			begin = end + 1
//			for begin < l{
//				if nums[begin] != 0{
//					break
//				}
//				begin++
//			}
//			end = begin
//			continue
//		}
//		if neg_cnt % 2 == 0{
//			max_length = int(math.Max(float64(max_length),float64(end - begin + 1)))
//		}
//		end++
//	}
//	return max_length
//}