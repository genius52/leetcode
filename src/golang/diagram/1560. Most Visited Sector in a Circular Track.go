package diagram

//Input: n = 4, rounds = [1,3,1,2]
//Output: [1,2]
func MostVisited(n int, rounds []int) []int{
	var l int = len(rounds)
	var start int = rounds[0]
	var end int = rounds[l - 1]
	var res []int
	if start <= end{
		for i := start;i <= end;i++{
			res = append(res,i)
		}
	}else{
		for i := 1;i <= end;i++{
			res = append(res,i)
		}
		for i := start;i <= n;i++{
			res = append(res,i)
		}
	}
	return res
}

//func MostVisited(n int, rounds []int) []int{
//	var record []int = make([]int,n + 1)
//	var l int = len(rounds)
//	record[rounds[0]]++
//	var max_cnt int = 0
//	for i := 1;i < l;i++{
//		start := rounds[i - 1]
//		end := rounds[i]
//		for j := start + 1;;j = (j + 1) % (n + 1){
//			idx := j % (n + 1)
//			record[idx]++
//			if record[idx] > max_cnt{
//				max_cnt = record[idx]
//			}
//			if j == end{
//				break
//			}
//		}
//	}
//	var res []int
//	for i := 1;i <= n;i++{
//		if record[i] == max_cnt{
//			res = append(res,i)
//		}
//	}
//	return res
//}

//func MostVisited(n int, rounds []int) []int {
//	var l int = len(rounds)
//	var record map[int]int = make(map[int]int)
//	var res []int
//	record[rounds[0]]++
//	var max_cnt int = 0
//	for i := 1;i < l;i++{
//		start := rounds[i - 1]
//		end := rounds[i]
//		if start > end{
//			end += n
//		}
//		for start < end{
//			index := start % n + 1
//			record[index]++
//			start++
//			if record[index] > max_cnt{
//				max_cnt = record[index]
//			}
//		}
//	}
//	for i := 1;i <= n;i++{
//		if cnt,ok := record[i];ok{
//			if cnt == max_cnt{
//				res = append(res,i)
//			}
//		}
//	}
//	return res
//}