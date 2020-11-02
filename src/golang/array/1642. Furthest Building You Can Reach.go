package array

func FurthestBuilding(heights []int, bricks int, ladders int) int {
	var l int = len(heights)
	if l <= 1{
		return 0
	}
	var index int = 0
	for index < l - 1{
		diff := heights[index + 1] - heights[index]
		if diff <= 0 {
			index++
			continue
		}
		if bricks >= diff{
			bricks -= diff
			index++
			continue
		}
		if ladders > 0{
			ladders--
			index++
			continue
		}
		break
	}
	return index
}