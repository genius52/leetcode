package number

func AddDigits(num int) int {
	for num > 9{
		var sum int = 0
		var cur int = num
		for cur > 0{
			mod := cur % 10
			sum += mod
			cur = cur /10
		}
		num = sum
	}
	return num
}