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

func LenLongestFibSubseq2(A []int) int{
	l := len(A)
	var dp [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		dp[i] = make([]int,l)
	}
	var res int = 0

	return res
}