package number

func SumEvenAfterQueries(A []int, queries [][]int) []int {
	var l int = len(queries)
	var res []int = make([]int,l)
	var even_sum int = 0
	for _,v := range A{
		if (v | 1) != v{
			even_sum += v
		}
	}
	for i := 0;i < l;i++{
		val := queries[i][0]
		idx := queries[i][1]
		if (A[idx] | 1) ==  A[idx]{//odd number
			if (val | 1) == val{
				even_sum += A[idx] + val
			}
		}else{//even number
			if (val | 1) == val{
				even_sum -= A[idx]
			}else{
				even_sum += val
			}
		}
		A[idx] += val
		res[i] = even_sum
	}
	return res
}