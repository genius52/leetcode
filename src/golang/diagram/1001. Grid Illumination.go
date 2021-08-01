package diagram

func GridIllumination(n int, lamps [][]int, queries [][]int) []int {
	var row_record map[int]int = make(map[int]int)
	var col_record map[int]int = make(map[int]int)
	var dia1_record map[int]int = make(map[int]int)
	var dia2_record map[int]int = make(map[int]int)
	var lamp_record map[int64]bool = make(map[int64]bool)
	for _,lamp := range lamps{
		r := lamp[0]
		c := lamp[1]
		if _,ok := lamp_record[int64(r * 1000000000 + c)];ok{
			continue
		}
		lamp_record[int64(r * 1000000000 + c)] = true
		if _,ok := row_record[r];!ok{
			row_record[r] = 1
		}else{
			row_record[r]++
		}
		if _,ok := col_record[c];!ok{
			col_record[c] = 1
		}else{
			col_record[c]++
		}
		if _,ok := dia1_record[r - c];!ok{
			dia1_record[r - c] = 1
		}else{
			dia1_record[r - c]++
		}
		if _,ok := dia2_record[r + c];!ok{
			dia2_record[r + c] = 1
		}else{
			dia2_record[r + c]++
		}

	}
	var dirs [][]int = [][]int{{-1,-1},{-1,0},{-1,1},{0,-1},{0,0},{0,1},{1,-1},{1,0},{1,1}}
	var l int = len(queries)
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		r := queries[i][0]
		c := queries[i][1]
		var light bool = false
		if _,ok := row_record[r];ok{
			light = true
		}
		if _,ok := col_record[c];ok{
			light = true
		}
		if _,ok := dia1_record[r - c];ok{
			light = true
		}
		if _,ok := dia2_record[r + c];ok{
			light = true
		}
		if light{
			res[i] = 1
		}
		for _,dir := range dirs{
			next_r := r + dir[0]
			next_c := c + dir[1]
			key := int64(next_r * 1000000000 + next_c)
			if _,ok := lamp_record[key];ok{
				delete(lamp_record,key)
				row_record[next_r]--
				if row_record[next_r] == 0{
					delete(row_record,next_r)
				}
				col_record[next_c]--
				if col_record[next_c] == 0{
					delete(col_record,next_c)
				}
				dia1_record[next_r - next_c]--
				if dia1_record[next_r - next_c] == 0{
					delete(dia1_record,next_r - next_c)
				}
				dia2_record[next_r + next_c]--
				if dia2_record[next_r + next_c] == 0{
					delete(dia2_record,next_r + next_c)
				}
			}
		}
	}
	return res
}