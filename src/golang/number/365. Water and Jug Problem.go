package number

func CanMeasureWater(x int, y int, z int) bool {
	if x == z || y == z{
		return true
	}
	var small_bottle,big_bottle int
	if x > y{
		small_bottle = y
		big_bottle = x
	}else{
		small_bottle = x
		big_bottle = y
	}
	if z > small_bottle + big_bottle{
		return false
	}
	var record map[int64]bool = make(map[int64]bool)
	var q [][]int
	q = append(q, []int{0,0})
	for len(q) != 0{
		ele := q[0]
		q = append(q[:0],q[1:]...)
		if ele[0] + ele[1] == z{
			return true
		}
		k := int64(ele[0]) * 1000000 + int64(ele[1])
		if _,ok := record[k];ok{
			continue
		}
		record[k] = true

		q = append(q,[]int{small_bottle,ele[1]})//灌满小杯子
		q = append(q,[]int{ele[0],big_bottle})//灌满大杯子
		//往大杯子里灌小杯水
		if ele[1] + small_bottle >= big_bottle{
			q = append(q,[]int{ele[0],big_bottle})
		}else{
			q = append(q,[]int{ele[0],ele[1] + small_bottle})
		}
		small_empty := small_bottle - ele[0]

		if ele[1] >= small_empty{
			q = append(q,[]int{small_bottle,ele[1] - small_empty})//把大杯子的水放入小杯子
		}else{
			q = append(q,[]int{ele[0] + ele[1],0})
		}
		big_empty := big_bottle - ele[1]
		if big_empty > 0{
			if ele[0] >= big_empty{
				q = append(q,[]int{ele[0] - big_empty,big_bottle})//把小杯子水放入大杯子
			}else{
				q = append(q,[]int{0,ele[1] + ele[0]})
			}
		}
		q = append(q,[]int{0,ele[0]})
	}
	return false
}