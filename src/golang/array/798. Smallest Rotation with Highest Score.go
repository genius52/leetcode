package array

//给定一个数组 A，我们可以将它按一个非负整数 K 进行轮调，这样可以使数组变为 A[K], A[K+1], A{K+2], ... A[A.length - 1], A[0], A[1], ..., A[K-1] 的形式。
//此后，任何值小于或等于其索引的项都可以记作一分。
type begin_end struct{
	begin int
	end int
}
//2, 3, 1, 4, 0
func BestRotation(A []int) int {
	var l int = len(A)
	var record []begin_end// = make([]Interval,l)//record[i] = 数字A[i]刚好不得分的偏移量
	for i := 0;i < l;i++{//每个数字
		if A[i] == 0{
			continue
		}
		if A[i] <= i{
			//A[5] = 2
			// goods.append([0, i-A[i]])
			//                if i+1<n:
			//                    goods.append([i+1, n-1])
			var cur1 begin_end
			cur1.begin = 0
			cur1.end = i - A[i]
			record = append(record, cur1)
			if i + 1 < l{
				var cur2 begin_end
				cur2.begin = i + 1
				cur2.end = l - 1
				record = append(record, cur2)
			}
		}else{
			//goods.append([i+1, i+n-A[i]])
			//A[2] = 5，l = 6
			//k = 1, A[1] = 5,不满足
			//k = 3,A[5] = 5,满足
			//k = 4,A[4] = 5
			//k = 5,A[3] = 5
			var cur begin_end
			cur.begin = i + 1
			cur.end = i + l - A[i]
			record = append(record, cur)
		}
	}
	var cnt []int = make([]int,l + 1)
	for _,r := range record{
		cnt[r.begin]++
		cnt[r.end + 1]--

	}
	var res int = 0
	var max_cnt int = 0
	var sum int = 0
	for i := 0;i < l;i++{
		sum += cnt[i]
		if sum > max_cnt{
			max_cnt = sum
			res = i
		}
	}
	return res
}



//TLE
//func BestRotation(A []int) int {
//	var l int = len(A)
//	var record []int = make([]int,l)//record[i] = 数字A[i]刚好不得分的偏移量
//	for i := 0;i < l;i++{//偏移量
//		for j := 0;j < l;j++{
//			index := (j + l - i) % l
//			if A[j] <= index{
//				record[i]++
//			}
//		}
//	}
//	var res int = 0
//	var max_scores int = 0
//	for i := 0;i < l;i++{
//		if record[i] > max_scores{
//			max_scores = record[i]
//			res = i
//		}
//	}
//	return res
//}