package array

func intervalIntersection(A [][]int, B [][]int) [][]int {
	var res [][]int
	index_a,index_b := 0,0
	for index_a < len(A) && index_b < len(B){
		if A[index_a][1] < B[index_b][0]{
			index_a++
		}else if B[index_b][1] < A[index_a][0]{
			index_b++
		}else{
			max_start := 0
			if A[index_a][0] > B[index_b][0]{
				max_start = A[index_a][0]
			}else{
				max_start = B[index_b][0]
			}
			min_end := 0
			if A[index_a][1] > B[index_b][1]{
				min_end = B[index_b][1]
				index_b++
			}else{
				min_end = A[index_a][1]
				index_a++
			}
			res = append(res, []int{max_start,min_end})
		}
	}
	return res
}