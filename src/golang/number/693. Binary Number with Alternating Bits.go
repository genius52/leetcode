package number

func hasAlternatingBits(n int) bool {
	var data []bool
	for n > 0{
		if (n | 1) == n{
			data = append(data,true)
		}else{
			data = append(data,false)
		}
		n >>= 1
	}
	var l int = len(data)
	if l <= 1{
		return true
	}
	for i := 1;i < l;i++{
		if data[i] == !data[i - 1]{
			continue
		}else{
			return false
		}
	}
	return true
}