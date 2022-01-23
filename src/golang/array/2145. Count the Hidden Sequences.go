package array

func numberOfArrays(differences []int, lower int, upper int) int {
	var l int = len(differences)
	var min_val int = 0
	var max_val int = 0
	var cur_val int = 0
	for i := 0;i < l;i++{
		cur_val += differences[i]
		if cur_val > max_val{
			max_val = cur_val
		}
		if cur_val < min_val{
			min_val = cur_val
		}
	}
	cur_diff := max_val - min_val
	target_diff := upper - lower
	if cur_diff > target_diff{
		return 0
	}
	return target_diff - cur_diff + 1
}