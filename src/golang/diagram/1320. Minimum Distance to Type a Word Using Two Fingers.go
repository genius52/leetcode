package diagram

func dfs_minimumDistance(word string,l int,pos int,row_f1 int,col_f1 int,row_f2 int,col_f2 int,memo *[]map[int]int)int{
	if pos >= l{
		return 0
	}
	key := row_f1 + col_f1 * 100 + row_f2 * 10000 + col_f2 * 1000000
	if row_f1 != -1 && row_f2 != -1 {
		if _,ok := (*memo)[pos][key];ok{
			return (*memo)[pos][key]
		}
	}
	row := int((word[pos] - 'A') / 6)
	col := int((word[pos] - 'A') % 6)
	//finger1 move to word[pos]
	var res1 int = 0
	if row_f1 == -1{
		res1 = dfs_minimumDistance(word,l,pos + 1,row,col,row_f2,col_f2,memo)
	}else{
		res1 = abs_int(row - row_f1) + abs_int(col - col_f1) + dfs_minimumDistance(word,l,pos + 1,row,col,row_f2,col_f2,memo)
	}
	var res2 int = 0
	//finger2 move to word[pos]
	if row_f2 == -1{
		res2 = dfs_minimumDistance(word,l,pos + 1,row_f1,col_f1,row,col,memo)
	}else{
		res2 = abs_int(row - row_f2) + abs_int(col - col_f2) + dfs_minimumDistance(word,l,pos + 1,row_f1,col_f1,row,col,memo)
	}
	(*memo)[pos][key] = min_int(res1,res2)
	return (*memo)[pos][key]
}

func MinimumDistance(word string) int {
	var l int = len(word)
	var memo []map[int]int = make([]map[int]int,l)
	for i := 0;i < l;i++{
		memo[i] = make(map[int]int)

	}
	return dfs_minimumDistance(word,l,0,-1,-1,-1,-1,&memo)
}