package array

import "sort"

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	var l int = len(asteroids)
	var cur int = mass
	for i := 0; i < l; i++ {
		if cur < asteroids[i] {
			return false
		}
		cur += asteroids[i]
	}
	return true
}
