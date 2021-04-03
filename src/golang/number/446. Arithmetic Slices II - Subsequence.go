package number

//Input: [2, 4, 6, 8, 10]
//
//Output: 7
//
//Explanation:
//All arithmetic subsequence slices are:
//[2,4,6]
//[4,6,8]
//[6,8,10]
//[2,4,6,8]
//[4,6,8,10]
//[2,4,6,8,10]
//[2,6,10]
func NumberOfArithmeticSlices(A []int) int {
	var l int = len(A)
	if l < 3{
		return 0
	}
	//dp[i][gap] means before index i(include) there are x elements their diff is gap
	var dp []map[int]int = make([]map[int]int,l)
	var res int = 0
	for i := 1;i < l;i++{
		dp[i] = make(map[int]int)
		for j := i - 1;j >= 0;j--{
			var diff int64 = int64(A[j]) - int64(A[i])
			if diff > 2147483647 || diff < -2147483648{
				continue
			}
			if _,ok := dp[i][int(diff)];!ok{
				dp[i][int(diff)] = 1
			}else{
				dp[i][int(diff)]++
			}
			if cnt,ok := dp[j][int(diff)];ok{
				res += cnt
				dp[i][int(diff)] += cnt
			}
		}
	}
	return res
}