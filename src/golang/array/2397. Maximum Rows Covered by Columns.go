package array

func MaximumRows(mat [][]int, cols int) int{
	var rows int = len(mat)
	var columns int = len(mat[0])
	if cols == columns{
		return rows
	}
	var res int = 0
	var limit int =  1 << columns
	for mask := 1;mask < limit;mask++{
		var ones int = 0
		var n int = mask
		for n > 0{
			if n | 1 == n{
				ones++
				if ones > cols{
					break
				}
			}
			n >>= 1
		}
		if ones != cols{
			continue
		}
		var cnt int = 0
		for i := 0;i < rows;i++{
			var choose bool = true
			for j := 0;j < columns;j++{
				if mat[i][j] == 1 && (mask | (1 << j) != mask){
					choose = false
					break
				}
			}
			if choose{
				cnt++
			}
		}
		if cnt > res{
			res = cnt
		}
	}
	return res
}

//func dfs_maximumRows(record map[int]int,cur_num int,offset int,cols int,max_num int)int{
//	if offset > 31{
//		return 0
//	}
//	if cur_num > max_num{
//		return 0
//	}
//	if cols == 0{
//		var res int = 0
//		for n,cnt := range record{
//			if (cur_num | n) == cur_num{
//				res += cnt
//			}
//		}
//		return res
//	}
//	//rest_bit := 32 - offset
//	//if rest_bit < cols{
//	//	return 0
//	//}
//	//set current bit as 1
//	next_num := cur_num | (1 << offset)
//	res1 := dfs_maximumRows(record,next_num,offset + 1,cols - 1,max_num)
//	//skip current bit
//	res2 := dfs_maximumRows(record,cur_num,offset + 1,cols,max_num)
//	return max_int(res1,res2)
//}
//
//func MaximumRows(mat [][]int, cols int) int {
//	var rows int = len(mat)
//	var columns int = len(mat[0])
//	if cols == columns{
//		return rows
//	}
//	var record map[int]int = make(map[int]int)
//	var max_one int = 0
//	for i := 0;i < rows;i++{
//		var cur int = 0
//		var one_cnt int = 0
//		for j := 0;j < columns;j++{
//			if mat[i][j] == 1{
//				one_cnt++
//				if one_cnt > cols{
//					break
//				}
//				cur |= 1 << j
//			}
//		}
//		if one_cnt > max_one{
//			max_one = one_cnt
//		}
//		if one_cnt <= cols{
//			record[cur]++
//		}
//	}
//	var max_num int = 1 << columns
//	return dfs_maximumRows(record,0,0,min_int(max_one,cols),max_num)
//}