package number

import "sort"

func MaximizeXor(nums []int, queries [][]int) []int {
	sort.Ints(nums)
	var idx int = 0
	var l int = len(queries)
	var queries2 []int = make([]int,l)//queries2[i] = 原本在queries的索引
	for i := 0;i < l;i++{
		queries2[i] = i
	}
	sort.Slice(queries2, func(i, j int) bool {
		return queries[queries2[i]][1] < queries[queries2[j]][1]
	})
	var root Trie
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		for idx < len(nums) && nums[idx] <= queries[queries2[i]][1]{
			var visit *Trie = &root
			for j := 31;j >= 0;j--{
				if nums[idx] | (1 << j) == nums[idx]{
					if visit.next_1 == nil{
						visit.next_1 = new(Trie)
					}
					visit = visit.next_1
				}else{
					if visit.next_0 == nil{
						visit.next_0 = new(Trie)
					}
					visit = visit.next_0
				}
			}
			visit.val = nums[idx]
			idx++
		}
		var find *Trie = &root
		if find.next_0 == nil && find.next_1 == nil{
			res[queries2[i]] = -1
			continue
		}
		for j := 31;j >= 0;j--{
			if (queries[queries2[i]][0] | (1 << j)) == queries[queries2[i]][0]{
				if find.next_0 != nil{
					find = find.next_0
				}else{
					find = find.next_1
				}
			}else{
				if find.next_1 != nil{
					find = find.next_1
				}else{
					find = find.next_0
				}
			}
		}
		res[queries2[i]] = queries[queries2[i]][0] ^ find.val
	}
	return res
}

//Plain tie tree
//func MaximizeXor(nums []int, queries [][]int) []int {
//	var root Ties
//	var min_num int = 2147483647
//	sort.Ints(nums)
//	for _,n := range nums{
//		if n < min_num{
//			min_num = n
//		}
//		var visit *Ties = &root
//		if n == 0{
//			if visit.next_zero == nil{
//				visit.next_zero = new(Ties)
//			}
//			visit.next_zero.is_num = true
//		}else{
//			for n > 0{
//				if n | 1 == n{
//					if visit.next_one == nil{
//						visit.next_one = new(Ties)
//					}
//					visit.next_one.val = 1
//					visit = visit.next_one
//				}else{
//					if visit.next_zero == nil{
//						visit.next_zero = new(Ties)
//					}
//					visit.next_zero.val = 0
//					visit = visit.next_zero
//				}
//				n = n >> 1
//			}
//			visit.is_num = true
//		}
//	}
//	var l int = len(queries)
//	var res []int = make([]int,l)
//	for i := 0;i < l;i++{
//		if queries[i][1] < min_num{
//			res[i] = -1
//			continue
//		}
//		var max_val int = -1
//		dfs_maximizeXor(queries[i][0],root.next_zero,0,0,queries[i][1],&max_val)
//		dfs_maximizeXor(queries[i][0],root.next_one,0,0,queries[i][1],&max_val)
//		res[i] = max_val
//	}
//	return res
//}