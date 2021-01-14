package array

//Input: arr = [1,2,3,10,4,2,3,5]
//Output: 3
//Explanation: The shortest subarray we can remove is [10,4,2] of length 3. The remaining elements after that will be [1,2,3,3,5] which are sorted.
//Another correct solution is to remove the subarray [3,10,4].
func FindLengthOfShortestSubarray(arr []int) int {
	var l int = len(arr)
	var left int = 1
	var right int = l - 2
	for left < l && arr[left] >= arr[left - 1]{
		left++
	}
	if left >= l{
		return 0
	}
	for right >= 0 && arr[right] <= arr[right + 1]{
		right--
	}
	var res int = min_int(l - left,right + 1)
	left--
	right++
	var i int = 0
	var j int = right
	for i <= left && j < l{
		if arr[i] <= arr[j]{
			res = min_int(res,j - i - 1)
			i++
		}else{
			j++
		}
	}
	return res
}