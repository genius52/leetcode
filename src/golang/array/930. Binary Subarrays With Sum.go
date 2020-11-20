package array

//Input: A = [1,0,1,0,1], S = 2
//Output: 4
//Explanation:
//The 4 subarrays are bolded below:
//[1,0,1,0,1]
//[1,0,1,0,1]
//[1,0,1,0,1]
//[1,0,1,0,1]
//A.length <= 30000
//0 <= S <= A.length
//A[i] is either 0 or 1.
func NumSubarraysWithSum(A []int, S int) int {
	var l int = len(A)
	var res int = 0
	var begin_origin int = 0
	var begin_visit int = 0
	var total int = 0
	for end := 0;end < l;end++{
		total += A[end]
		if total < S{
			continue
		}
		for begin_origin <= end && total > S{
			total -= A[begin_origin]
			begin_origin++
		}
		begin_visit = begin_origin
		for begin_visit <= end && total == S{
			res++
			if A[begin_visit] == 0{
				begin_visit++
			}else{
				begin_visit = begin_origin
				break
			}
		}
	}
	return res
}