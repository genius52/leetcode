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
func NumSubarraysWithSum2(A []int, S int) int {
	var l int = len(A)
	var record map[int]int = make(map[int]int)
	var res int = 0
	var total int = 0
	for i := 0;i < l;i++{
		total += A[i]
		if cnt,ok := record[total - S];ok{
			res += cnt
		}
		if total == S{
			res++
		}
		record[total]++
	}
	return res
}

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