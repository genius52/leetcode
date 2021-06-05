package diagram

import "math"

func escapeGhosts(ghosts [][]int, target []int) bool {
	var dis int = int(math.Abs(float64(target[0])) + math.Abs(float64(target[1])))
	for _,pos := range ghosts{
		cur := int(math.Abs(float64(pos[0] - target[0])) + math.Abs(float64(pos[1] - target[1])))
		if cur <= dis{
			return false
		}
	}
	return true
}