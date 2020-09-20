package array

//Input: arr = [1,4,2,5,3]
//Output: 58
//Explanation: The odd-length subarrays of arr and their sums are:
//[1] = 1
//[4] = 4
//[2] = 2
//[5] = 5
//[3] = 3
//[1,4,2] = 7
//[4,2,5] = 11
//[2,5,3] = 10
//[1,4,2,5,3] = 15
func SumOddLengthSubarrays(arr []int) int {
	var l int = len(arr)
	var res int = 0
	for length := 1;length <= l;length += 2{//length
		for j := 0;j < l;j++{//start index
			for k := 0;k < length && (j + length) <= l;k++{//0 - length
				res += arr[j + k]
			}
		}
	}
	return res
}