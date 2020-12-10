package array

//Input: customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
//Output: 16
//Explanation: The bookstore owner keeps themselves not grumpy for the last 3 minutes.
//The maximum number of customers that can be satisfied = 1 + 1 + 1 + 1 + 7 + 5 = 16.
func MaxSatisfied(customers []int, grumpy []int, X int) int{
	var window int = 0
	var total int = 0
	var l int = len(customers)
	for i := 0;i < l && i < X;i++{
		if grumpy[i] == 0{
			total += customers[i]
		}
		if grumpy[i] == 1{
			window += customers[i]
		}
	}
	var max_val int = window
	for i := X;i < l;i++{
		if grumpy[i] == 0{
			total += customers[i]
		}
		if grumpy[i - X] == 1{
			window -= customers[i - X]
		}
		if grumpy[i] == 1{
			window += customers[i]
		}
		if window > max_val{
			max_val = window
		}
	}
	return max_val + total
}

//func dp_maxSatisfied(customers []int,start int,end int,grumpy []int)int{
//	var l int = len(customers)
//	var total int = 0
//	for i := 0;i < l;i++{
//		if i >= start && i < end{
//			total += customers[i]
//		}else{
//			if grumpy[i] == 0{
//				total += customers[i]
//			}
//		}
//	}
//	return total
//}
////TLE
//func MaxSatisfied(customers []int, grumpy []int, X int) int {
//	var l int = len(customers)
//	var res int = 0
//	for i := 0;i + X - 1 < l;i++{
//		//if customers[i] == 0 || grumpy[i] == 0{
//		//	continue
//		//}
//		cur := dp_maxSatisfied(customers,i,i + X,grumpy)
//		if cur > res{
//			res = cur
//		}
//	}
//	return res
//}