package number

//left = 1, right = 22
//Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]
func check(num int)bool{
	var n int = num
	for n > 0{
		mod := n % 10
		if mod == 0{
			return false
		}
		if (num % mod) != 0{
			return false
		}
		n = n / 10
	}
	return true
}

func SelfDividingNumbers(left int, right int) []int {
	var res []int
	for i := left;i <= right;i++{
		if i < 10{
			res = append(res,i)
		}else{
			if check(i){
				res = append(res,i)
			}
		}
	}
	return res
}