package number

//Input: n = 10
//Output: 16
func dp_getMoneyAmount(low int,high int,memo [][]int)int{
	if low >= high{
		return 0
	}
	if low == (high - 1){
		return low
	}
	if memo[low][high] != 0{
		return memo[low][high]
	}
	var min_cost int = 2147483647
	for i := low;i <= high;i++{
		//guess number is too large
		cost1 := i + dp_getMoneyAmount(low,i - 1,memo)

		//guess number is too small
		cost2 := i + dp_getMoneyAmount(i+1, high,memo)
		cur := max_int(cost1,cost2)
		if cur < min_cost{
			min_cost = cur
		}
	}
	memo[low][high] = min_cost
	return min_cost
}

func GetMoneyAmount(n int) int {
	if n <= 1{
		return 0
	}
	var min_cost int = 2147483647
	var memo [][]int = make([][]int,n + 1)
	for i := 0;i <= n;i++{
		memo[i] = make([]int,n + 1)
	}
	for i := 1;i <= n;i++{
			//guess number is too large
		cost1 := i + dp_getMoneyAmount(1,i-1,memo)
		//guess number is too small
		cost2 := i + dp_getMoneyAmount(i+1,n,memo)
		cur := max_int(cost1,cost2)
		if cur < min_cost{
			min_cost = cur
		}
	}
	return min_cost
}