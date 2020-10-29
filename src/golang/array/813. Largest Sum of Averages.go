package array

//func dp_largestSumOfAverages(A []int,pos int,select_nums []int,rest_groups int,memo [][]float64)float64{
//	if memo[rest_groups][pos] != 0{
//		return memo[rest_groups][pos]
//	}
//	if rest_groups == 0{
//		var total float64 = 0
//		var l int = 0
//		for _,n := range select_nums{
//			total += float64(n)
//			l++
//		}
//		for i := pos;i < len(A);i++{
//			total += float64(A[i])
//			l++
//		}
//		return total/float64(l)
//	}
//	var total float64 = 0
//	for _,n := range select_nums{
//		total += float64(n)
//	}
//	var skip_cur = total/float64(len(select_nums)) +  dp_largestSumOfAverages(A,pos + 1,[]int{A[pos]},rest_groups - 1,memo)
//
//	rest_num := len(A) - pos - 1
//	var add_cur float64 = 0.0
//	if rest_groups <= rest_num{
//		select_nums = append(select_nums,A[pos])
//		add_cur = dp_largestSumOfAverages(A,pos + 1,select_nums,rest_groups,memo)
//	}
//	memo[rest_groups][pos] = math.Max(skip_cur,add_cur)
//	return memo[rest_groups][pos]
//}

//func LargestSumOfAverages(A []int, K int) float64 {
//	var memo [][]float64 = make([][]float64,len(A) + 1)
//	for i := 0;i <= len(A);i++{
//		memo[i] = make([]float64,len(A) + 1)
//	}
//	return dp_largestSumOfAverages(A,1,[]int{A[0]},K - 1,memo)
//}

func dp_largestSumOfAverages(A []int,start int,rest_groups int,memo [][]float64)float64{
	if rest_groups < 0{
		return 0
	}
	var l int = len(A)
	if rest_groups == 0{
		var res float64 = 0
		for i := start;i < l ;i++ {
			res += float64(A[i])
		}
		return res / float64(l - start)
	}
	if memo[start][rest_groups] > 0{
		return memo[start][rest_groups]
	}

	var res float64 = 0
	var sub_sum float64 = 0
	for end := start;end + rest_groups < l ;end++{
		sub_sum += float64(A[end])
		next_sum := dp_largestSumOfAverages(A,end + 1,rest_groups - 1,memo)
		var total float64 = sub_sum / float64(end - start + 1) + next_sum
		if total > res{
			res = total
		}
	}
	memo[start][rest_groups] = res
	return res
}

func LargestSumOfAverages(A []int, K int) float64 {
	var memo [][]float64 = make([][]float64,len(A) + 1)
	for i := 0;i <= len(A);i++{
		memo[i] = make([]float64,len(A) + 1)
	}
	return dp_largestSumOfAverages(A,0,K - 1,memo)
}