package number

import "sort"

func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
	sort.Ints(enemyEnergies)
	var l int = len(enemyEnergies)
	var score int64 = 0
	var right int = l - 1
	if currentEnergy >= enemyEnergies[0] {
		score += int64(currentEnergy / enemyEnergies[0])
		currentEnergy -= int(score) * enemyEnergies[0]
	}
	for score > 0 && right > 0 {
		currentEnergy += enemyEnergies[right]
		right--
	}
	if currentEnergy >= enemyEnergies[0] {
		score += int64(currentEnergy / enemyEnergies[0])
	}
	return score
}
