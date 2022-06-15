package array

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	var l int = len(spells)
	var l2 int = len(potions)
	var res []int = make([]int,l)
	sort.Ints(potions)
	for i := 0;i < l;i++{
		target := success/int64(spells[i])
		if success % int64(spells[i]) != 0{
			target++
		}
		idx := sort.Search(l2, func(i int) bool {
			return potions[i] >= int(target)
		})
		res[i] = l2 - idx
	}
	return res
}