package array

func stableMountains(height []int, threshold int) []int {
	var res []int
	var find bool = false
	for i := 0; i < len(height); i++ {
		if find {
			res = append(res, i)
		}
		if height[i] > threshold {
			find = true
		} else {
			find = false
		}
	}
	return res
}
