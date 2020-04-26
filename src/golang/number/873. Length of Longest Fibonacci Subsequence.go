package number

//873
//Input: [1,2,3,4,5,6,7,8]
//Output: 5
//Explanation:
//The longest subsequence that is fibonacci-like: [1,2,3,5,8].
//Input: [1,3,7,11,12,14,18]
//Output: 3
//Explanation:
//The longest subsequence that is fibonacci-like:
//[1,11,12], [3,11,14] or [7,11,18].
func LenLongestFibSubseq(A []int) int{
	l := len(A)
	if l <= 2{
		return 0
	}
	var res int = 0
	var record map[int]bool = make(map[int]bool)
	for i := 0;i < l;i++{
		record[A[i]] = true
	}
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			first := A[i]
			second := A[j]
			cnt := 2
			_,ok := record[first + second]
			for ok{
				first,second = second,first + second
				cnt++
				_,ok = record[first + second]
			}
			res = max_int(cnt,res)
		}
	}
	if res <= 2{
		return 0
	}
	return res
}

//dp[i][j] means the length which last 2 element is index i and index j
func LenLongestFibSubseq2(A []int) int{
	l := len(A)
	var dp [][]int = make([][]int,l)
	var record map[int]int = make(map[int]int)
	for i := 0;i < l;i++{
		dp[i] = make([]int,l)
		record[A[i]] = i
	}
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			dp[i][j] = 2
		}
	}
//[2,4,7,8,9,10,14,15,18,23,32,50]
	var res int = 0
	for i := 2;i < l;i++{
		for j := i - 1;j > 0;j--{
			first_num := A[i] - A[j]
			if _,ok := record[first_num];ok{
				pos := record[first_num]
				if first_num >= A[j]{
					break
				}
				dp[j][i] = 1 + dp[pos][j]
				if dp[j][i] > res{
					res = dp[j][i]
				}
			}else{
				dp[j][i] = 2
			}
		}
	}

	return res
}