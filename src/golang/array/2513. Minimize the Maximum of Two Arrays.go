package array

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func check_minimizeSet(divisor1 int, divisor2 int,uniqueCnt1 int, uniqueCnt2 int,max_num int)bool{
	//不能被d1且不能被d2整除的数的数量
	cnt1 := max_num - max_num/divisor1
	if cnt1 < uniqueCnt1{
		return false
	}
	cnt2 := max_num - max_num/divisor2
	if cnt2 < uniqueCnt2{
		return false
	}
	lcm_cnt := max_num / lcm(divisor1,divisor2)
	union_cnt :=  max_num - lcm_cnt
	if union_cnt < (uniqueCnt1 + uniqueCnt2){
		return false
	}
	return true
}

func MinimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	var low int = 1
	var high int = (uniqueCnt1 + uniqueCnt2) * 2
	//var res int = 1e9
	for low < high{
		var mid int = (low + high)/2
		ret := check_minimizeSet(divisor1,divisor2,uniqueCnt1,uniqueCnt2,mid)
		if ret{
			//res = mid
			high = mid
		}else{
			low = mid + 1
		}
	}
	return low
}