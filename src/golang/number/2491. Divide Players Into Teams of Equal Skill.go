package number

import "sort"

func dividePlayers(skill []int) int64 {
	var l int = len(skill)
	sort.Ints(skill)
	var target int = skill[0] + skill[l - 1]
	//var MOD int64 = 1e9 + 7
	var res int64 = int64(skill[0] * skill[l - 1])
	for i := 1;i < l/2;i++{
		if skill[i] + skill[l - 1 - i] != target{
			return -1
		}
		res += int64(skill[i] * skill[l - 1 - i])
		//res %= MOD
	}
	return res
}