package number

func dfs_minimumOperationsToMakeEqual(x int, target int, memo []int) int {
	if x == target {
		return 0
	}
	if x < target {
		return target - x
	}
	if memo[x] != 1<<31-1 {
		return memo[x]
	}
	var res int = x - target
	if x%11 == 0 {
		res = min_int(res, 1+dfs_minimumOperationsToMakeEqual(x/11, target, memo))
	} else {
		mod := x % 11
		if x-mod >= 11 {
			res = min_int(res, mod+1+dfs_minimumOperationsToMakeEqual((x-mod)/11, target, memo))
		}
		res = min_int(res, 11-mod+1+dfs_minimumOperationsToMakeEqual((x+11-mod)/11, target, memo))
	}
	if x%5 == 0 {
		res = min_int(res, 1+dfs_minimumOperationsToMakeEqual(x/5, target, memo))
	} else {
		mod := x % 5
		if x-mod >= 5 {
			res = min_int(res, mod+1+dfs_minimumOperationsToMakeEqual((x-mod)/5, target, memo))
		}
		res = min_int(res, 5-mod+1+dfs_minimumOperationsToMakeEqual((x+5-mod)/5, target, memo))
	}
	memo[x] = res
	return res
}

func MinimumOperationsToMakeEqual(x int, y int) int {
	var memo []int = make([]int, 100001)
	for i := 0; i <= 100000; i++ {
		memo[i] = 1<<31 - 1
	}
	return dfs_minimumOperationsToMakeEqual(x, y, memo)
}

//func minimumOperationsToMakeEqual(x int, y int) int {
//	var res int = abs_int(x - y)
//	var memo []int = make([]int, x+1)
//	for i := 0; i <= x; i++ {
//		memo[i] = 1<<31 - 1
//	}
//	type val_steps struct {
//		val   int
//		steps int
//	}
//	var q list.List
//	q.PushBack(val_steps{y, 0})
//	var cur_step int = 0
//	for q.Len() > 0 {
//		var cur_len int = q.Len()
//		for i := 0; i < cur_len; i++ {
//			var cur val_steps = q.Front().Value.(val_steps)
//			q.Remove(q.Front())
//			if cur.val < x {
//				//plus one
//
//				//* 5
//
//				//* 11
//			} else {
//				memo[cur_step] = min_int(memo[cur_step], cur.steps+cur.val-x)
//			}
//		}
//		cur_step++
//	}
//	return res
//}
