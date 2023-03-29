package number

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	var res int = 0
	if k >= numOnes {
		k -= numOnes
		res += numOnes
	} else {
		res += k
		k = 0
	}

	if k >= numZeros {
		k -= numZeros
	} else {
		k = 0
	}
	return res - k
}
