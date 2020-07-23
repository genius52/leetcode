package number

import (
	"math"
	"sort"
)

func NumMovesStones(a int, b int, c int) []int {
	var pos []int = []int{a,b,c}
	sort.Ints(pos)
	if(pos[0] == pos[1] - 1 && pos[1] == pos[2] - 1){
		return []int{0,0}
	}
	var res []int
	var min_gap int = int(math.Min(float64(pos[1] - pos[0]), float64(pos[2] - pos[1])))
	if(min_gap == 1 || min_gap == 2){
		res = append(res,1)
	}else{
		res = append(res,2)
	}
	res = append(res,pos[2] - pos[0] - 2)
	return res
}
