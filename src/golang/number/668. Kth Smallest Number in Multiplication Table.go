package number

//Input: m = 3, n = 3, k = 5
//Output: 3
//Explanation: The 5th smallest number is 3.
func count_lessthan(m int,n int,target int)int{
	var cnt int = 0
	for i := 1;i <= m;i++{
		cnt += min_int(n,target/i)
	}
	return cnt
}

func FindKthNumber2(m int, n int, k int) int {
	var low int = 1
	var high int = m * n
	for low < high{
		mid := (low + high)/2
		cnt := count_lessthan(m,n,mid)
		if cnt < k{
			low = mid + 1
		}else{
			high = mid
		}
	}
	return low
}