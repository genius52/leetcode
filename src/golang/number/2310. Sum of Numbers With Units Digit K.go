package number


func MinimumNumbers(num int, k int) int {
	if num == 0{
		return 0
	}
	last := num % 10
	if k == 0{
		if last == 0{
			return 1
		}else{
			return -1
		}
	}
	var res int = 2147483647
	for i := 0;i <= 9;i++{
		cur := i * 10 + last
		if cur != 0 && cur <= num{
			mod := cur % k
			if mod == 0{
				cnt := cur / k
				if cnt < res{
					res = cnt
				}
			}
		}
	}
	if res == 2147483647{
		return -1
	}
	return res
}