package number

//Input: n = 6
//Output: true
//Explanation: 6 = 2 × 3
func IsUgly(n int) bool {
	for n > 0{
		mod := n % 2
		if mod != 0{
			break
		}
		n /= 2
	}
	for n > 0{
		mod := n % 3
		if mod != 0{
			break
		}
		n /= 3
	}
	for n > 0{
		mod := n % 5
		if mod != 0{
			break
		}
		n /= 5
	}
	if n == 1{
		return true
	}else{
		return false
	}
}