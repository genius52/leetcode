package array

//Input: arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
//Output: 3
//Explanation: Sub-arrays [2,5,5],[5,5,5] and [5,5,8] have averages 4, 5 and 6 respectively.
//All other sub-arrays of size 3 have averages less than 4 (the threshold).
func numOfSubarrays(arr []int, k int, threshold int) int {
	l := len(arr)
	if l < k {
		return 0
	}
	var target int = threshold * k
	var sum int = 0
	for i := 0;i < k;i++{
		sum += arr[i]
	}
	var res int = 0
	begin := 0
	end := begin + k - 1
	for end < l{
		if sum >= target{
			res++
		}
		end++
		if end < l{
			sum -= arr[begin]
			sum += arr[end]
		}
		begin++
	}
	return res
}