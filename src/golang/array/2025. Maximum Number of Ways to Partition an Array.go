package array

func WaysToPartition(nums []int, k int) int {
	var l int = len(nums)
	var prefix []int = make([]int,l)
	var suffix []int = make([]int,l)
	prefix[0] = nums[0]
	for i := 1;i < l;i++{
		prefix[i] = prefix[i - 1] + nums[i]
	}
	suffix[l - 1] = nums[l - 1]
	for i := l - 2;i >= 0;i--{
		suffix[i] = suffix[i + 1] + nums[i]
	}
	var bigger_than_suffix map[int][]int = make(map[int][]int)
	var less_than_suffix map[int][]int = make(map[int][]int)
	var res int = 0
	for i := 0;i < l - 1;i++{
		if prefix[i] == suffix[i + 1]{
			res++
		}else{
			cur_diff := abs_int(suffix[i + 1] - prefix[i])
			if prefix[i] > suffix[i + 1]{
				bigger_than_suffix[cur_diff] = append(bigger_than_suffix[cur_diff],i)
			}else{
				less_than_suffix[cur_diff] = append(less_than_suffix[cur_diff],i)
			}
		}
	}
	for i := 0;i < l;i++{
		//old_diff := prefix[i] - suffix[i + 1]
		//new_diff := k - nums[i] + prefix[i] - suffix[i + 1]
		diff := abs_int(k - nums[i])//把nums[i]改成k后,前缀数组增加了increase_diff
		var cnt int = 0
		if k > nums[i]{//替换了更大的数字, prefix[i] 到 prefix[l - 1]变大了，suffix[i] 到 suffix[0]变大了
			if idx,ok := less_than_suffix[diff];ok{
				for j := len(idx) - 1;j >= 0;j--{
					if less_than_suffix[diff][j] < i{
						break
					}
					cnt++
				}
			}
			if idx,ok := bigger_than_suffix[diff];ok{
				for j := 0;j < len(idx);j++{
					if bigger_than_suffix[diff][j] >= i{
						break
					}
					cnt++
				}
			}
		}else{//替换了更小的数字, prefix[i] 到 prefix[l - 1]变小了，suffix[i] 到 suffix[0]变小了
			if idx,ok := bigger_than_suffix[diff];ok{
				for j := len(idx) - 1;j >= 0;j--{
					if bigger_than_suffix[diff][j] < i{
						break
					}
					cnt++
				}
			}
			if idx,ok := less_than_suffix[diff];ok{
				for j := 0;j < len(idx);j++{
					if less_than_suffix[diff][j] >= i{
						break
					}
					cnt++
				}
			}
		}
		res = max_int(res,cnt)
	}
	return res
}