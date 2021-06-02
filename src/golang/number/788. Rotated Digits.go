package number

func is_good(N int)bool{
	var rotate_exist bool = false
	for N > 0{
		mod := N % 10
		if mod == 3 || mod == 4 || mod == 7{
			return false
		}
		if mod == 2 ||  mod == 5 || mod == 6 || mod == 9{
			rotate_exist = true
		}
		N /= 10
	}
	return rotate_exist
}

func rotatedDigits(N int) int {
	var res int = 0
	for i := 1;i <= N;i++{
		if is_good(i){
			res++
		}
	}
	return res
}