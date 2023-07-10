package array

import "sort"

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	var l int = len(nums)
	var stones map[int]bool = make(map[int]bool)
	for i := 0; i < l; i++ {
		stones[nums[i]] = true
	}
	for i := 0; i < len(moveTo); i++ {
		if _, ok := stones[moveFrom[i]]; ok {
			delete(stones, moveFrom[i])
			stones[moveTo[i]] = true
		}
	}
	var res []int
	for k, _ := range stones {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}

//func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
//	var l int = len(nums)
//	var src_dst map[int]int = make(map[int]int)     //from to
//	var dst_src map[int][]int = make(map[int][]int) //to from
//	for i := 0; i < len(moveTo); i++ {
//		//1 --> 2,2 --> 3
//		//1 --> 3
//		if pre_srcs, ok := dst_src[moveFrom[i]]; ok {
//			for _, pre_src := range pre_srcs {
//				src_dst[pre_src] = moveTo[i]
//			}
//		}
//		src_dst[moveFrom[i]] = moveTo[i]
//		dst_src[moveTo[i]] = append(dst_src[moveTo[i]], moveFrom[i])
//	}
//	var record map[int]bool = make(map[int]bool)
//	for i := 0; i < l; i++ {
//		if pos, ok := src_dst[nums[i]]; ok {
//			record[pos] = true
//		} else {
//			record[nums[i]] = true
//		}
//	}
//	var res []int
//	for k, _ := range record {
//		res = append(res, k)
//	}
//	sort.Ints(res)
//	return res
//}
