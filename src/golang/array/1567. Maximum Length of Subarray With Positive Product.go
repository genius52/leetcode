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