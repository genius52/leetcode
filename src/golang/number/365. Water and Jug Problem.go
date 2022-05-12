package number

func dfs_CanMeasureWater(target int,choose []int)bool{
	if(target == 0){
		return true
	}
	if(target < 0){
		return false
	}
	for _,n := range choose{
		ret := dfs_CanMeasureWater(target - n,choose)
		if(ret){
			return true
		}
	}
	return false
}

func CanMeasureWater(x int, y int, z int) bool {
	if(x == z || y == z){
		return true
	}
	var small_bottle,big_bottle int
	if(x > y){
		small_bottle = y
		big_bottle = x
	}else{
		small_bottle = x
		big_bottle = y
	}
	if(z > small_bottle + big_bottle){
		return false
	}
	var record map[int64]bool = make(map[int64]bool)
	var q [][]int
	q = append(q, []int{0,0})
	record[0] = true
	for len(q) != 0{
		ele := q[0]
		q = append(q[:0],q[1:]...)
		if ele[0] + ele[1] == z{
			return true
		}
		var possible [][]int
		possible = append(possible,[]int{small_bottle,ele[1]})//灌满小杯子
		possible = append(possible,[]int{ele[0],big_bottle})//灌满大杯子
		//往大杯子里灌小杯水
		if ele[1] + small_bottle >= big_bottle{
			possible = append(possible,[]int{ele[0],big_bottle})
		}else{
			possible = append(possible,[]int{ele[0],ele[1] + small_bottle})
		}
		small_empty := small_bottle - ele[0]

		if(ele[1] >= small_empty){
			possible = append(possible,[]int{small_bottle,ele[1] - small_empty})//把大杯子的水放入小杯子
		}else{
			possible = append(possible,[]int{ele[0] + ele[1],0})
		}
		big_empty := big_bottle - ele[1]
		if(big_empty > 0){
			if(ele[0] >= big_empty){
				possible = append(possible,[]int{ele[0] - big_empty,big_bottle})//把小杯子水放入大杯子
			}else{
				possible = append(possible,[]int{0,ele[1] + ele[0]})
			}
		}
		possible = append(possible,[]int{0,ele[0]})
		for _, p := range possible{
			k := int64(p[0]) * 1000000 + int64(p[1])
			if _,ok := record[k];!ok{
				record[k] = true
				q = append(q,p)
			}
		}
	}
	return false
}