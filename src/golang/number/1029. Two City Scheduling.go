package number

import "sort"

//[[518,518],[71,971],[121,862],[967,607],[138,754],[513,337],[499,873],[337,387],[647,917],[76,417]]
//

type A_B [][]int

func (s A_B) Less(i, j int) bool {
	if (s[i][0] - s[i][1]) <= (s[j][0] - s[j][1]){
		return true
	}else {
		return false
	}
}

func (s A_B) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s A_B) Len() int {
	return len(s)
}

func twoCitySchedCost(costs [][]int) int{
	sort.Sort(A_B(costs))
	var res int = 0
	var l int = len(costs)
	for i := 0;i < l/2;i++{
		res += costs[i][0]
	}
	for i := l/2;i < l;i++{
		res += costs[i][1]
	}
	return res
}

//func twoCitySchedCost(costs [][]int) int {
//	var res int = 0
//	var num int = len(costs)
//	if num <= 1{
//		return res
//	}
//	var a_cnt int = 0
//	var b_cnt int = 0
//	for i := 0;i < num;i++{
//		for j := 1;j < num - i;j++{
//			if math.Abs(float64(costs[j][0] - costs[j][1])) > math.Abs(float64(costs[j-1][0] - costs[j-1][1])){
//				costs[j],costs[j-1] = costs[j-1],costs[j]
//			}
//		}
//	}
//	for _,ele := range costs{
//		if ele[0] > ele[1] && b_cnt < num/2 {
//			res += ele[1]
//			b_cnt++
//		} else{
//			res += ele[0]
//			a_cnt++
//		}
//	}
//	return res
//}