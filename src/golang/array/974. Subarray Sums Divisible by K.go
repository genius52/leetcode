package array

//Input: A = [4,5,0,-2,-3,1], K = 5
//Output: 7
//Explanation: There are 7 subarrays with a sum divisible by K = 5:
//[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
func SubarraysDivByK(A []int, K int) int {
	var l int = len(A)
	var record map[int]int = make(map[int]int)
	var total int = 0
	var res int = 0
	for i := 0;i < l;i++{
		total += A[i]
		rest := total % K
		if rest < 0{
			rest += K
		}
		if rest == 0{
			res++
		}
		if cnt,ok := record[rest];ok{
			res += cnt
			record[rest]++
		}else{
			record[rest] = 1
		}
	}
	return res
}