package array

//Input: A = [3,6,9,12]
//Output: 4
//Explanation:
//The whole array is an arithmetic sequence with steps of length = 3.
//Input: A = [9,4,7,2,10]
//Output: 3
//Explanation:
//The longest arithmetic subsequence is [4,7,10].
func LongestArithSeqLength(A []int) int {
	var l int = len(A)
	var record []map[int]int = make([]map[int]int,l)
	var max_cnt int = 2
	for i := 0;i < l;i++{
		record[i] = make(map[int]int)
		for j := i - 1;j >= 0 ;j--{
			diff := A[i] - A[j]
			if cnt,ok := record[j][diff];ok{
				if cnt_i,ok2 := record[i][diff];ok2{
					if cnt + 1 > cnt_i{
						record[i][diff] = cnt + 1
					}
				}else{
					record[i][diff] = cnt + 1
				}
				if record[i][diff] > max_cnt{
					max_cnt = record[i][diff]
				}
			}else{
				if _,ok2 := record[i][diff];!ok2{
					record[i][diff] = 2
				}
				if record[i][diff] > max_cnt{
					max_cnt = record[i][diff]
				}
			}
		}
	}
	return max_cnt
}