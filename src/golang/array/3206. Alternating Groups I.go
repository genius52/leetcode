package array

func numberOfAlternatingGroups(colors []int) int {
	var l int = len(colors)
	var res int = 0
	for i := 0; i < l; i++ {
		left := (i - 1 + l) % l
		right := (i + 1) % l
		if colors[i] != colors[left] && colors[i] != colors[right] {
			res++
		}
	}
	return res
}
