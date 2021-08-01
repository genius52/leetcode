package array

import "math"

//Input: A = [2,1,2,4,2,2], B = [5,2,6,2,3,2]
//Output: 2
//Explanation:
//The first figure represents the dominoes as given by A and B: before we do any rotations.
//If we rotate the second and fourth dominoes, we can make every value in the top row equal to 2, as indicated by the second figure.
func MinDominoRotations2(A []int, B []int) int{
	var countA []int = make([]int,7)
	var countB []int = make([]int,7)
	var same_cnt []int = make([]int,7)
	var l int = len(A)
	for i := 0;i < l;i++{
		if A[i] == B[i]{
			same_cnt[A[i]]++
		}
		countA[A[i]]++
		countB[B[i]]++
	}
	var res int = 2147483647
	for i := 1;i <= 6;i++{
		if countA[i] + countB[i] - same_cnt[i] < l{
			continue
		}
		swaps := min_int(countA[i] - same_cnt[i],countB[i] - same_cnt[i])
		if swaps < res{
			res = swaps
		}
	}
	if res == 2147483647{
		return -1
	}
	return res
}

func MinDominoRotations(A []int, B []int) int {
	var l int = len(A)
	if l <= 1{
		return 0
	}
	var count []int = make([]int,7)
	var max_len int = 0
	var max_len_num int = 0
	for i := 0;i < l;i++{
		if A[i] == B[i]{
			count[A[i]]++
			if count[A[i]] > max_len{
				max_len_num = A[i]
				max_len = count[A[i]]
			}
		}else{
			count[A[i]]++
			count[B[i]]++
			if count[A[i]] > max_len{
				max_len_num = A[i]
				max_len = count[A[i]]
			}
			if count[B[i]] > max_len{
				max_len_num = B[i]
				max_len = count[B[i]]
			}
		}
		if max_len != i + 1{
			return -1
		}
	}
	if max_len != l{
		return -1
	}
	var swap_a int = 0
	var swap_b int = 0
	for i := 0;i < l;i++{
		if A[i] != max_len_num{
			swap_a++
		}
		if B[i] != max_len_num{
			swap_b++
		}
	}
	return int(math.Min(float64(swap_a),float64(swap_b)))
}