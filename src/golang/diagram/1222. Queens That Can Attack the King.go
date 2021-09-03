package diagram

func recursive_queensAttacktheKing(record map[int]bool,r int,c int,dir []int,res *[][]int){
	if r < 0 || r >= 8 || c < 0 || c >= 8{
		return
	}
	if _,ok := record[r * 10 + c];ok{
		*res = append(*res,[]int{r,c})
		return
	}
	recursive_queensAttacktheKing(record,r + dir[0],c + dir[1],dir,res)
}

func queensAttacktheKing(queens [][]int, king []int) [][]int{
	var dirs [][]int = [][]int{{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}}
	var res [][]int
	var record map[int]bool = make(map[int]bool)
	for _,queen := range queens{
		record[queen[0] * 10 + queen[1]] = true
	}
	for _,dir := range dirs{
		recursive_queensAttacktheKing(record,king[0] + dir[0],king[1] + dir[1],dir,&res)
	}
	return res
}

//func check_point_inline(start []int,end []int,third []int) bool{
//
//	if (third[0] - start[0])*(end[1] - start[1]) == (end[0] - start[0]) * (third[1] - start[1]) &&
//		math.Min(float64(start[0]),float64(end[0])) < float64(third[0]) && float64(third[0]) < math.Max(float64(start[0]),float64(end[0])) &&
//		math.Min(float64(start[1]),float64(end[1])) < float64(third[1]) && float64(third[1]) < math.Max(float64(start[1]),float64(end[1])){
//		return true
//	}else{
//		return false;
//	}
//}
//
//func queensAttacktheKing(queens [][]int, king []int) [][]int {
//	var res [][] int
//	var samelines [][]int
//	var samecolumns [][]int
//	var sametilts [][]int
//	for i := 0; i < len(queens);i++{
//		if queens[i][0] == king[0]{
//			samelines = append(samelines, queens[i])
//		}else if queens[i][1] == king[1]{
//			samecolumns = append(samecolumns,queens[i])
//		}else if math.Abs(float64(queens[i][0] - king[0])) == math.Abs(float64(queens[i][1] - king[1])){
//			sametilts = append(sametilts,queens[i])
//		}
//	}
//	min_left_distance := -1
//	min_right_distance := 8
//	var left_pos []int
//	var right_pos []int
//	for i := 0;i < len(samelines);i++{
//		if samelines[i][1] < king[1]{
//			if samelines[i][1] > min_left_distance{
//				min_left_distance = samelines[i][1]
//				left_pos = samelines[i]
//			}
//		}else{
//			if samelines[i][1] < min_right_distance{
//				min_right_distance = samelines[i][1]
//				right_pos = samelines[i]
//			}
//		}
//	}
//	if len(left_pos) > 0{
//		res = append(res, left_pos)
//	}
//	if len(right_pos) > 0{
//		res = append(res,right_pos)
//	}
//	min_upper_pos := -1
//	min_bottom_pos := 8
//	var upper_pos []int
//	var bottom_pos []int
//	for i := 0;i < len(samecolumns) ;i++ {
//		if samecolumns[i][0] < king[0]{
//			if samecolumns[i][0] > min_upper_pos{
//				min_upper_pos = samecolumns[i][0]
//				upper_pos = samecolumns[i]
//			}
//		}else{
//			if samecolumns[i][0] < min_bottom_pos{
//				min_bottom_pos = samecolumns[i][0]
//				bottom_pos = samecolumns[i]
//			}
//		}
//	}
//	if len(upper_pos) > 0{
//		res = append(res, upper_pos)
//	}
//	if len(bottom_pos) > 0{
//		res = append(res,bottom_pos)
//	}
//	var min_tilt_collection [][]int //save nearest queens
//	var first bool = true
//	for i := 0;i < len(sametilts);i++{
//		if first{
//			first = false
//			min_tilt_collection = append(min_tilt_collection, sametilts[i])
//			continue
//		}
//		var add_to_result bool = true
//		for j := 0;j < len(min_tilt_collection);{
//			if check_point_inline(king,sametilts[i],min_tilt_collection[j]) {
//				add_to_result = false
//				break
//			}
//
//			if check_point_inline(king,min_tilt_collection[j],sametilts[i]) {
//				fmt.Println("tilt ", sametilts[i], " closer than ", min_tilt_collection[j])
//				min_tilt_collection = append(min_tilt_collection[:j], min_tilt_collection[j+1:]...)
//			}else{
//				j++
//			}
//		}
//		if add_to_result{
//			fmt.Println( sametilts[i], "  is valid")
//			min_tilt_collection = append(min_tilt_collection,sametilts[i])
//		}
//	}
//	res = append(res,min_tilt_collection...)
//	return res
//}