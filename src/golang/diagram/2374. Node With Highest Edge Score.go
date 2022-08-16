package diagram

func edgeScore(edges []int) int {
	var l int = len(edges)
	var score []int = make([]int,l)
	for i := 0;i < l;i++{
		score[edges[i]] += i
	}
	var max_score int = 0
	var res int = 0
	for i := 0;i < l;i++{
		if score[i] > max_score{
			max_score = score[i]
			res = i
		}
	}
	return res
}