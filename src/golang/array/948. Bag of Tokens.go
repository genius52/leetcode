package array

import (
	"sort"
)

func BagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	var l int = len(tokens)
	var res int = 0
	var begin int = 0
	var end int = l - 1
	var scores int = 0
	var cur_power int = P
	for begin <= end{
		if cur_power >= tokens[begin]{
			cur_power -= tokens[begin]
			scores++
			begin++
			if scores > res{
				res = scores
			}
		}else{
			if scores > 0{
				cur_power += tokens[end]
				scores--
				end--
			}else{
				break
			}
		}
	}
	return res
}