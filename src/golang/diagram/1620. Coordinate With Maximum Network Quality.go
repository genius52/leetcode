package diagram

import (
	"math"
	"sort"
)

func BestCoordinate(towers [][]int, radius int) []int {
	var l int = len(towers)
	var max_val float64 = 0
	var idx [][]int
	for i := 0;i < 51;i++{
		for j := 0;j < 51;j++{
			var total float64 = 0
			for k := 0;k < l;k++{
				dis := math.Sqrt(float64((i - towers[k][0]) * (i - towers[k][0]) +
					(j - towers[k][1]) * (j - towers[k][1])))
				if dis <= float64(radius){
					total += math.Floor(float64(towers[k][2]) / float64(1 + dis))
				}
			}
			if total == 0{
				continue
			}
			if total > max_val{
				max_val = total
				idx = [][]int{[]int{i,j}}
			}else if total == max_val{
				idx = append(idx,[]int{i,j})
			}
		}
	}
	if max_val == 0{
		return []int{0,0}
	}
	sort.Slice(idx, func(i, j int) bool {
		if idx[i][0] != idx[j][0]{
			return idx[i][0] < idx[j][0]
		}else{
			return idx[i][1] < idx[j][1]
		}
	})
	return []int{idx[0][0],idx[0][1]}
}