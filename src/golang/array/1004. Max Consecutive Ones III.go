package array

//Input: A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
//Output: 6
//Explanation:
//[1,1,1,0,0,1,1,1,1,1,1]

//Input: A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
//Output: 10
//Explanation:
//[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
func LongestOnes(A []int, K int) int {
	var l int = len(A)
	var begin int = 0
	var end int = 0
	var zero_cnt int = 0
	var max_len int = 0
	for begin < l && end < l{
		for end < l && zero_cnt <= K{
			if A[end] == 0{
				zero_cnt++
			}
			end++
		}
		cur_len := end - begin - 1
		if end == l && zero_cnt <= K{
			cur_len++
		}
		if cur_len > max_len{
			max_len = cur_len
		}
		for begin < l && zero_cnt > K{
			if A[begin] == 0{
				zero_cnt--
			}
			begin++
		}
	}
	return max_len
}