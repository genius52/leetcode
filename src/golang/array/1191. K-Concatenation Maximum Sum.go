package array

//Input: arr = [1,2], k = 3
//Output: 9
//Example 2:
//
//Input: arr = [1,-2,1], k = 5
//Output: 2
//Example 3:
//
//Input: arr = [-1,-2], k = 7
//Output: 0
func KConcatenationMaxSum2(arr []int, k int) int{
	var l int = len(arr)
	var pre int = arr[0]
	var max_sub int = arr[0]
	var sum int = arr[0]
	for i := 1;i < l;i++{
		sum += arr[i]
		sum = sum % 1000000007
		cur := max_int(pre + arr[i],arr[i])
		pre = cur
		if cur > max_sub{
			max_sub = cur
		}
	}
	if max_sub <= 0{
		return 0
	}
	if k == 1{
		return max_sub
	}
	var prefix int = 0
	var suffix int = 0
	var max_prefix int = 0
	var max_suffix int = 0
	for i := 0;i < l;i++{
		if i == 0{
			prefix = arr[0]
			suffix = arr[l - 1]
		}else{
			prefix += arr[i]
			suffix += arr[l - 1 - i]
		}
		prefix = prefix % 1000000007
		suffix = suffix % 1000000007
		max_prefix = max_int(max_prefix,prefix)
		max_suffix = max_int(max_suffix,suffix)
	}
	if sum < 0{
		return max_int_number(max_sub % 1000000007,(max_prefix % 1000000007 + max_suffix % 1000000007 )% 1000000007)
	}else{
		var res1 int64 = int64(max_sub)
		var res2 int64 = (int64(k) - 2) * int64(sum) + int64(max_prefix) + int64(max_suffix)
		if res1 > res2{
			return max_sub % 1000000007
		}else{
			return (((k - 2) * sum) % 1000000007 + max_prefix % 1000000007 + max_suffix % 1000000007) % 1000000007
		}
	}
}

func KConcatenationMaxSum(arr []int, k int) int {
	var l int = len(arr)
	var dp_small []int = make([]int,l)
	dp_small[0] = arr[0]
	var dp_big []int = make([]int,l)
	dp_big[0] = arr[0]
	var total int = arr[0]
	var min_sum int = arr[0]
	var max_sum int = arr[0]
	var max_prefix int = arr[0]
	var max_suffix int = arr[l - 1]
	var reverse_total int = arr[l - 1]
	for i := 1;i < l;i++{
		total += arr[i]
		if total > max_prefix{
			max_prefix = total
		}
		reverse_total += arr[l - 1 - i]
		if reverse_total > max_suffix{
			max_suffix = reverse_total
		}
		dp_small[i] = min_int(dp_small[i - 1] + arr[i],arr[i])
		if dp_small[i] < min_sum{
			min_sum = dp_small[i]
		}
		dp_big[i] = max_int(dp_big[i - 1] + arr[i],arr[i])
		if dp_big[i] > max_sum{
			max_sum = dp_big[i]
		}
	}
	var max_circle_sum int = total - min_sum
	if k == 1{
		return max_sum
	}else{
		if total <= 0{
			return max_int(max_circle_sum,max_sum)
		}else{
			var res int = (k - 2) * total % 1000000007 + max_suffix + max_prefix
			return res
		}
	}
}