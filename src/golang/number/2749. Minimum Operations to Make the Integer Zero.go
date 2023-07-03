package number

func check_number(n int, target_cnt int) bool {
	if n == target_cnt {
		return true
	}
	if n < 0 {
		return false
	}
	var n2 int = n
	var one_cnt int = 0
	var right_zero_cnt int = 0
	for n2 > 0 {
		if n2|1 == n2 {
			one_cnt++
		} else {
			right_zero_cnt++
		}
		n2 >>= 1
	}
	if one_cnt > target_cnt {
		return false
	}
	if one_cnt == target_cnt {
		return true
	}
	//if (one_cnt + right_zero_cnt) >= target_cnt {
	//	return true
	//}
	return false
}

func MakeTheIntegerZero(num1 int, num2 int) int {
	if num1 < num2 {
		return -1
	}
	var cnt int = 1
	for cnt <= 60 {
		var n int = num1
		n -= num2 * cnt
		if check_number(n, cnt) {
			return cnt
		}
		cnt++
	}
	if cnt > 60 {
		return -1
	}
	return cnt
}
