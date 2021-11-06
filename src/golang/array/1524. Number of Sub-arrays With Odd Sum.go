package array

//Input: arr = [1,3,5]
//Output: 4
//Explanation: All sub-arrays are [[1],[1,3],[1,3,5],[3],[3,5],[5]]
//All sub-arrays sum are [1,4,9,3,8,5].
//Odd sums are [1,9,3,5] so the answer is 4.
func NumOfSubarrays2(arr []int) int{
	var pre_odd_cnt int = 0
	var pre_even_cnt = 0
	var res int = 0
	if arr[0] | 1 == arr[0]{
		pre_odd_cnt = 1
		res = 1
	}else{
		pre_even_cnt = 1
	}
	var l int = len(arr)
	for i := 1;i < l;i++{
		if (arr[i] | 1) == arr[i]{//当前是奇数
			cur_odd_cnt := pre_even_cnt + 1 //还要加上自己
			res += cur_odd_cnt
			cur_even_cnt := pre_odd_cnt
			pre_odd_cnt = cur_odd_cnt
			pre_even_cnt = cur_even_cnt
		}else{//当前是偶数
			res += pre_odd_cnt
			pre_even_cnt++
		}
		res = res % (1e9 + 7)
	}
	return res
}

func NumOfSubarrays(arr []int) int {
	var l int = len(arr)
	var dp_odd []int = make([]int,l)//dp[i] = end with i odd arr num
	var dp_even []int = make([]int,l)//dp[i] = end with i even arr num
	if arr[0] | 1 == arr[0]{
		dp_odd[0] = 1
	}else{
		dp_even[0] = 1
	}
	for i := 1;i < l;i++{
		if arr[i] | 1 == arr[i]{
			dp_odd[i] = dp_even[i - 1] + 1
			dp_odd[i] %= 1000000007
			dp_even[i] = dp_odd[i - 1]
			dp_even[i] %= 1000000007
		}else{
			dp_even[i] = dp_even[i - 1] + 1
			dp_even[i] %= 1000000007
			dp_odd[i] = dp_odd[i - 1]
			dp_odd[i] %= 1000000007
		}
	}
	var res int = 0
	for i := 0;i < l;i++{
		res += dp_odd[i]
		res %= 1000000007
	}
	return res
}

//数学方法计算排列数(从n中取m个数)
//func mathPailie(n int, m int) int {
//	return jieCheng(n) / jieCheng(n-m)
//}

//数学方法计算组合数(从n中取m个数)
func mathZuhe(n int, m int) int {
	return jieCheng(n) / (jieCheng(n-m) * jieCheng(m))
}

//阶乘
func jieCheng(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
