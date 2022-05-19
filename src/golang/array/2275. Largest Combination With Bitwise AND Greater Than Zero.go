package array

func largestCombination(candidates []int) int {
	var l int = len(candidates)
	var res int = 0
	for i := 0;i < 32;i++{
		mask := 1 << i
		var cnt int = 0
		for j := 0;j < l;j++{
			if (candidates[j] & mask) != 0{
				cnt++
			}
		}
		if cnt > res{
			res = cnt
		}
	}
	return res
}