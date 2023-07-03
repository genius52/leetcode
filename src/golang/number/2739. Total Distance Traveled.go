package number

func distanceTraveled(mainTank int, additionalTank int) int {
	var res int = 0
	for mainTank >= 5 && additionalTank > 0 {
		res += 5
		mainTank -= 5
		mainTank += 1
		additionalTank--
	}
	res += mainTank
	return res * 10
}
