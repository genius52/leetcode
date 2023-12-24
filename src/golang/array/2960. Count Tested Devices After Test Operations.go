package array

func countTestedDevices(batteryPercentages []int) int {
	var l int = len(batteryPercentages)
	var res int = 0
	for i := 0; i < l; i++ {
		if batteryPercentages[i] == 0 {
			continue
		}
		if batteryPercentages[i] <= res {
			continue
		}
		res++
	}
	return res
}
