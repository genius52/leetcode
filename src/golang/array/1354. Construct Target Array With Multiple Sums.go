package array

//func dfs_isPossible(cur []int,l int,sum int,target *[]int,target_sum int)bool{
//	if sum == target_sum{
//		return true
//	}
//	var cur_sum int = sum
//	for i := 0;i < l;i++{
//		if cur_sum > (*target)[i]{
//			continue
//		}
//		old := cur[i]
//		cur[i] = cur_sum
//		res := dfs_isPossible(cur,l,cur_sum + cur_sum - old,target,target_sum)
//		if res{
//			return true
//		}
//		cur[i] = old
//	}
//	return false
//}

//TLE
//func IsPossible(target []int) bool {
//	var l int = len(target)
//	if l == 1{
//		return target[0] == 1
//	}
//	var target_sum int = 0
//	for i := 0;i < l;i++{
//		target_sum += target[i]
//	}
//	var cur []int = make([]int,l)
//	for i := 0;i < l;i++{
//		cur[i] = 1
//	}
//	return dfs_isPossible(cur,l,l,&target,target_sum)
//}