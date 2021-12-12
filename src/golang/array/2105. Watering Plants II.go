package array

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	var l int = len(plants)
	var left int = 0
	var right int = l - 1
	var cur_a int = capacityA
	var cur_b int = capacityB
	var cnt int = 0
	for left < right {
		if cur_a < plants[left] {
			cur_a = capacityA - plants[left]
			cnt++
		} else {
			cur_a -= plants[left]
		}
		if cur_b < plants[right] {
			cur_b = capacityB - plants[right]
			cnt++
		} else {
			cur_b -= plants[right]
		}
		left++
		right--
	}
	if l|1 == l {
		if cur_a < plants[left] && cur_b < plants[right] {
			cnt++
		}
	}
	return cnt
}
