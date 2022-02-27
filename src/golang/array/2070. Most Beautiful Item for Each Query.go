package array

import "sort"

//Input: items = [[1,2],[3,2],[2,4],[5,6],[3,5]], queries = [1,2,3,4,5,6]
//Output: [2,4,5,5,6,6]
//Input: items = [[1,2],[1,2],[1,3],[1,4]], queries = [1]
//Output: [4]

//给你一个二维整数数组 items ，其中 items[i] = [pricei, beautyi] 分别表示每一个物品的 价格 和 美丽值 。
//同时给你一个下标从 0开始的整数数组queries。对于每个查询queries[j]，
//你想求出价格小于等于queries[j]的物品中，最大的美丽值是多少。如果不存在符合条件的物品，那么查询的结果为0
func MaximumBeauty(items [][]int, queries []int) []int{
	var l int = len(queries)
	var record [][2]int = make([][2]int,l)
	for i := 0;i < l;i++{
		record[i] = [2]int{queries[i],i}
	}
	//价格从低到高
	sort.Slice(record, func(i, j int) bool {
		return record[i][0] < record[j][0]
	})
	//价格从低到高
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	var res []int = make([]int,l)
	var max_beauty int = 0
	var j int = 0
	var l2 int = len(items)
	for i := 0;i < l;i++{
		for j < l2 && items[j][0] <= record[i][0]{
			max_beauty = max_int(max_beauty,items[j][1])
			j++
		}
		res[record[i][1]] = max_beauty
	}
	return res
}

func maximumBeauty(items [][]int, queries []int) []int {
	var l int = len(queries)
	var res []int = make([]int,l)
	sort.Slice(items, func(i, j int) bool {
		if items[i][1] != items[j][1]{
			return items[i][1] > items[j][1]
		}else{
			return items[i][0] < items[j][0]
		}
	})
	for i := 0;i < l;i++{
		var find bool = false
		for j := 0;j < len(items);j++{
			if items[j][0] <= queries[i]{
				res[i] = items[j][1]
				find = true
				break
			}
		}
		if !find{
			res[i] = 0
		}
	}
	return res
}