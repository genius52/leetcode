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