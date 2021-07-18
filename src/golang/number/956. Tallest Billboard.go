package number

func TallestBillboard(rods []int) int{
	return 0
}

//TLE
//func dp_tallestBillboard(rods []int,l int,pos int,l1 int,l2 int,sum int,memo []map[int]map[int]int)int{
//	if pos >= l{
//		if l1 == l2{
//			return l1
//		}else{
//			return 0
//		}
//	}
//	if l1 > sum/2 || l2 > sum/2{
//		return 0
//	}
//
//	if _,ok1 := memo[pos][l1];ok1{
//		if _,ok2 := memo[pos][l1][l2];ok2{
//			return memo[pos][l1][l2]
//		}else{
//			//memo[pos][l1] = make(map[int]int)
//		}
//	}else{
//		memo[pos] = make(map[int]map[int]int)
//		memo[pos][l1] = make(map[int]int)
//	}
//	var res int = 0
//	//skip
//	res = max_int(res,dp_tallestBillboard(rods,l,pos + 1,l1,l2,sum,memo))
//	//add cur to l1
//	res = max_int(res,dp_tallestBillboard(rods,l,pos + 1,l1 + rods[pos],l2,sum,memo))
//	//add cur to l2
//	res = max_int(res,dp_tallestBillboard(rods,l,pos + 1,l1,l2 + rods[pos],sum,memo))
//	memo[pos][l1][l2] = res
//	return res
//}
//
//func TallestBillboard(rods []int) int {
//	var l int = len(rods)
//	var sum int = 0
//	for i := 0;i < l;i++{
//		sum += rods[i]
//	}
//	var memo []map[int]map[int]int = make([]map[int]map[int]int,l)
//	return dp_tallestBillboard(rods,l,0,0,0,sum,memo)
//}