package array

func MergeTriplets(triplets [][]int, target []int) bool {
	var l int = len(triplets)
	var res []int = make([]int,3)
	for i := 0;i < l;i++{
		if triplets[i][0] > target[0] || triplets[i][1] > target[1] || triplets[i][2] > target[2]{
			continue
		}
		res[0] = max_int(res[0],triplets[i][0])
		res[1] = max_int(res[1],triplets[i][1])
		res[2] = max_int(res[2],triplets[i][2])
	}
	if res[0] == target[0] && res[1] == target[1] && res[2] == target[2]{
		return true
	}
	return false
}