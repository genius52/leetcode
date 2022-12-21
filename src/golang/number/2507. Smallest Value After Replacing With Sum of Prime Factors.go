package number

func SmallestValue(n int) int {
	for true{
		var n2 int = n
		var i int = 2
		var sum int = 0
		for n2 != 1{
			for n2 % i == 0{
				sum += i
				n2 /= i
				if n2 == 1{
					break
				}
			}
			i++
		}
		if sum >= n{
			return n
		}else{
			n = sum
		}
	}
	return n
}