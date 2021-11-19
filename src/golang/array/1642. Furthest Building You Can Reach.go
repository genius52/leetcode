package array

//MLE
//func dp_furthestBuilding(heights []int,l int,pos int,bricks int,ladders int,memo  map[int]map[int]map[int]int)int{
//	if pos == l - 1{
//		return pos
//	}
//	if memo[pos] == nil{
//		memo[pos] = make(map[int]map[int]int)
//	}
//	if memo[pos][bricks] == nil{
//		memo[pos][bricks] = make(map[int]int)
//	}
//	if memo[pos][bricks][ladders] != 0{
//		return memo[pos][bricks][ladders]
//	}
//	var res int = pos
//	if heights[pos + 1] <= heights[pos]{
//		res = dp_furthestBuilding(heights,l,pos + 1,bricks,ladders,memo)
//	}else{
//		diff := heights[pos + 1] - heights[pos]
//		if bricks >= diff{
//			res = dp_furthestBuilding(heights,l,pos + 1,bricks - diff,ladders,memo)
//		}
//		if ladders > 0{
//			res = max_int(res,dp_furthestBuilding(heights,l,pos + 1,bricks,ladders - 1,memo))
//		}
//	}
//	memo[pos][bricks][ladders] = pos
//	return res
//}
//
//func FurthestBuilding(heights []int, bricks int, ladders int) int {
//	var l int = len(heights)
//	var memo map[int]map[int]map[int]int = make(map[int]map[int]map[int]int)
//	return dp_furthestBuilding(heights,l,0,bricks,ladders,memo)
//}

//Wrong solution
//func FurthestBuilding(heights []int, bricks int, ladders int) int {
//	var l int = len(heights)
//	if l <= 1{
//		return 0
//	}
//	var index int = 0
//	for index < l - 1{
//		diff := heights[index + 1] - heights[index]
//		if diff <= 0 {
//			index++
//			continue
//		}
//		if bricks >= diff{
//			bricks -= diff
//			index++
//			continue
//		}
//		if ladders > 0{
//			ladders--
//			index++
//			continue
//		}
//		break
//	}
//	return index
//}