package number

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	var l1 int = len(players)
	var l2 int = len(trainers)
	var res int = 0
	var find_idx int = 0
	for i := 0;i < l1 && find_idx < l2;i++{
		for find_idx < l2 && players[i] > trainers[find_idx]{
			find_idx++
		}
		if find_idx < l2{
			res++
			find_idx++
		}
	}
	return res
}