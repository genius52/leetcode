package number

func maxBottlesDrunk(numBottles int, numExchange int) int {
	var res int = 0
	for numBottles >= numExchange {
		res += numExchange
		numBottles -= numExchange
		numBottles++
		numExchange++
	}
	res += numBottles
	return res
}
