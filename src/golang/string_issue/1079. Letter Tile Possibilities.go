package string_issue

func dfs_numTilePossibilities(record [26]int)int {
	var res int = 0
	for i := 0;i < 26;i++{
		if record[i] == 0{
			continue
		}
		res++//end with record[i]
		record[i]--//use current char so minus one
		res += dfs_numTilePossibilities(record) //add postfix count
		record[i]++
	}
	return res
}

func numTilePossibilities(tiles string) int {
	var record [26]int
	for _,t := range tiles{
		record[int(t - 'A')]++
	}
	return dfs_numTilePossibilities(record)
}