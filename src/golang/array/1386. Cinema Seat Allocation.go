package array

func MaxNumberOfFamilies(n int, reservedSeats [][]int) int{
	var record map[int]int = make(map[int]int)
	for _,seat := range reservedSeats{
		record[seat[0]] = record[seat[0]] | (1 << seat[1])
	}
	var res int = (n - len(record)) * 2
	var left int = 1 << 6 | 1 << 7 | 1 << 8 | 1 << 9
	var mid int = 1 << 4 | 1 << 5 | 1 << 6 | 1 << 7
	var right int = 1 << 2 | 1 << 3 | 1 << 4 | 1 << 5
	for _,v := range record{
		ret := (v & mid)
		if  ret == 0{
			if (v & left) == 0 &&  (v & right) == 0{
				res += 2
			}else{
				//if (record[i] & left) == 0 || (record[i] & right) == 0{
					res++
				//}
			}
		}else{
			if (v & left) == 0{
				res++
			}
			if (v & right) == 0{
				res++
			}
		}
	}
	return res
}

//TLE
//func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
//	var res int = 0
//	var block map[int][]bool = make(map[int][]bool)
//	for _,seat := range reservedSeats{
//		if _,ok := block[seat[0]];!ok{
//			block[seat[0]] = make([]bool,11)
//		}
//		block[seat[0]][seat[1]] = true
//	}
//	for i := 1;i <= n;i++{
//		if _,ok := block[i];!ok{
//			res += 2
//		}else{
//			if !block[i][4] && !block[i][5] && !block[i][6] && !block[i][7]{
//				if !block[i][2] && !block[i][3] && !block[i][8] && !block[i][9]{
//					res += 2
//				}else{
//					res++
//				}
//			}else{
//				if !block[i][2] && !block[i][3] && !block[i][4] && !block[i][5]{
//					res++
//				}
//				if !block[i][6] && !block[i][7] && !block[i][8] && !block[i][9]{
//					res++
//				}
//			}
//		}
//	}
//	return res
//}