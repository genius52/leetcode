package array

import "sort"

//[9,11,99,101]
//[[10,1],[7,1],[14,1],[100,1],[96,1],[103,1]]
//func dp_minimumTotalDistance(robot []int,l1 int,idx1 int,positions []int,l2 int,idx2 int,memo [][]int64)int64{
//	if idx1 == l1{
//		return 0
//	}
//	if idx2 == l2{
//		return 1 << 62 - 1
//	}
//	if memo[idx1][idx2] != -1{
//		return memo[idx1][idx2]
//	}
//	var res int64 = -1
//	ret1 := abs_int64(int64(robot[idx1]) - int64(positions[idx2])) + dp_minimumTotalDistance(robot,l1,idx1 + 1,positions,l2,idx2 + 1,memo)
//	ret2 := dp_minimumTotalDistance(robot,l1,idx1,positions,l2,idx2 + 1,memo)
//	if ret1 != -1 && ret2 != -1{
//		res = min_int64_number(ret1,ret2)
//	}else if ret1 == -1{
//		res = ret2
//	}else if ret2 == -1{
//		res = ret1
//	}
//	memo[idx1][idx2] = res
//	return res
//}
//
//func MinimumTotalDistance(robot []int, factory [][]int) int64 {
//	sort.Ints(robot)
//	sort.Slice(factory, func(i, j int) bool {
//		return factory[i][0] < factory[j][0]
//	})
//	var l1 int = len(robot)
//	var l2 int = len(factory)
//	var positions []int
//	for i := 0;i < l2;i++{
//		for j := 0;j < factory[i][1];j++{
//			positions = append(positions,factory[i][0])
//		}
//	}
//	var l3 int = len(positions)
//	var memo [][]int64 = make([][]int64,l1)
//	for i := 0;i < l1;i++{
//		memo[i] = make([]int64,l3)
//		for j := 0;j < l3;j++{
//			memo[i][j] = -1
//		}
//	}
//	return dp_minimumTotalDistance(robot,l1,0,positions,l3,0,memo)
//}

func dp_minimumTotalDistance(robot []int,l1 int,idx1 int,positions []int,l2 int,idx2 int,memo [][]int64)int64{
	if idx1 == l1{
		return 0
	}
	if idx2 == l2{
		return 1 << 62 - 1
	}
	if memo[idx1][idx2] != -1{
		return memo[idx1][idx2]
	}
	var res int64 = -1
	ret1 := abs_int64(int64(robot[idx1]) - int64(positions[idx2])) + dp_minimumTotalDistance(robot,l1,idx1 + 1,positions,l2,idx2 + 1,memo)
	ret2 := dp_minimumTotalDistance(robot,l1,idx1,positions,l2,idx2 + 1,memo)
	if ret1 != -1 && ret2 != -1{
		res = min_int64_number(ret1,ret2)
	}else if ret1 == -1{
		res = ret2
	}else if ret2 == -1{
		res = ret1
	}
	memo[idx1][idx2] = res
	return res
}

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})
	var l1 int = len(robot)
	var l2 int = len(factory)
	var positions []int
	for i := 0;i < l2;i++{
		for j := 0;j < factory[i][1];j++{
			positions = append(positions,factory[i][0])
		}
	}
	var l3 int = len(positions)
	var memo [][]int64 = make([][]int64,l1)
	for i := 0;i < l1;i++{
		memo[i] = make([]int64,l3)
		for j := 0;j < l3;j++{
			memo[i][j] = -1
		}
	}
	return dp_minimumTotalDistance(robot,l1,0,positions,l3,0,memo)
}