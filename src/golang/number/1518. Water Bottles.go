package number

func numWaterBottles(numBottles int, numExchange int) int {
	var res int = numBottles
	var empty int = numBottles
	for empty >= numExchange{
		res += empty / numExchange
		rest := empty % numExchange + empty / numExchange
		empty = rest
	}
	return res
}