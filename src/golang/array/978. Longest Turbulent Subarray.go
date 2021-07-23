package array


//For i <= k < j:
//arr[k] > arr[k + 1] when k is odd, and
//arr[k] < arr[k + 1] when k is even.
//Or, for i <= k < j:
//arr[k] > arr[k + 1] when k is even, and
//arr[k] < arr[k + 1] when k is odd.

//Input: arr = [9,4,2,10,7,8,8,1,9]
//Output: 5
//Explanation: arr[1] > arr[2] < arr[3] > arr[4] < arr[5]

//arr = [4,8,12,16]
func maxTurbulenceSize(arr []int) int {
	var l int = len(arr)
	var smaller_than_pre []int = make([]int,l)
	smaller_than_pre[0] = 1
	var bigger_than_pre []int = make([]int,l)
	bigger_than_pre[0] = 1
	var res int = 1
	for i := 1;i < l;i++{
		if arr[i] > arr[i - 1]{// 比前一个数大
			bigger_than_pre[i] = 1 + smaller_than_pre[i - 1]
			smaller_than_pre[i] = 1
		}else if arr[i] < arr[i - 1]{// 比前一个数小
			smaller_than_pre[i] = 1 + bigger_than_pre[i - 1]
			bigger_than_pre[i] = 1
		}else{
			smaller_than_pre[i] = 1
			bigger_than_pre[i] = 1
		}
		if smaller_than_pre[i] > res{
			res = smaller_than_pre[i]
		}
		if bigger_than_pre[i] > res{
			res = bigger_than_pre[i]
		}
	}
	return res
}