package array

import "sort"

//Input: mat =
//[[1,1,0,0,0],
//[1,1,1,1,0],
//[1,0,0,0,0],
//[1,1,0,0,0],
//[1,1,1,1,1]],
//k = 3
//Output: [2,0,3]
func KWeakestRows(mat [][]int, k int) []int{
	var rows int = len(mat)
	var columns int = len(mat[0])
	var record [][2]int ////record[i][0]: one cnt,record[i][1]:idx
	for i := 0;i < rows;i++{
		var one int = 0
		for j := 0;j < columns;j++{
			if mat[i][j] == 1{
				one++
			}
		}
		if len(record) < k{
			record = append(record,[2]int{one,i})
		}else{
			if one < record[k - 1][0]{
				record[k - 1][0] = one
				record[k - 1][1] = i
			}
		}
		sort.Slice(record, func(i, j int) bool {
			if record[i][0] == record[j][0]{
				return record[i][1] < record[j][1]
			}
			return record[i][0] < record[j][0]
		})
	}
	var res []int = make([]int,k)
	for i := 0;i < k;i++{
		res[i] = record[i][1]
	}
	return res
}

//func kWeakestRows(mat [][]int, k int) []int {
//	var rows = len(mat)
//	if rows == 0{
//		return []int{}
//	}
//	var columns = len(mat[0])
//	var record map[int][]int = make(map[int][]int)
//	for i := 0;i < rows;i++{
//		j := 0
//		for ;j < columns;j++ {
//			if mat[i][j] == 0 {
//				break
//			}
//		}
//		if _,ok := record[j];ok{
//			record[j] = append(record[j], i)
//		}else{
//			record[j] = []int{i}
//		}
//	}
//	var keys []int
//	for l := range record {
//		keys = append(keys, l)
//	}
//	sort.Ints(keys)
//	var res []int
//	var i int = 0
//	for _,num := range keys{
//		for _,row := range record[num]{
//			res = append(res,row)
//			i++
//			if i == k{
//				return res
//			}
//		}
//	}
//	return res
//}