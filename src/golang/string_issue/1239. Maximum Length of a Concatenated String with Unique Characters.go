package string_issue

func dfs_maxLength(arr []string,l int,cur_pos int,record [26]int)int{
	if cur_pos >= l{
		return 0
	}
	var sub_len int = len(arr[cur_pos])
	var dup bool = false
	var self_dup bool = false
	var cur_record [26]int
	for i := 0;i < sub_len;i++{
		if record[arr[cur_pos][i] - 'a'] > 0{
			dup = true
			break
		}
		cur_record[arr[cur_pos][i] - 'a']++
		if cur_record[arr[cur_pos][i] - 'a'] > 1{
			self_dup = true
		}
	}
	if dup || self_dup{
		return dfs_maxLength(arr,l,cur_pos + 1,record)
	}else{
		for i := 0;i < sub_len;i++{
			record[arr[cur_pos][i] - 'a']++
		}
		add_cur := sub_len + dfs_maxLength(arr,l,cur_pos + 1,record)
		for i := 0;i < sub_len;i++{
			record[arr[cur_pos][i] - 'a']--
		}
		skip_cur := dfs_maxLength(arr,l,cur_pos + 1,record)
		if add_cur > skip_cur{
			return add_cur
		}
		return skip_cur
	}
}

func MaxLength(arr []string) int {
	var l int = len(arr)
	var record [26]int
	return dfs_maxLength(arr,l,0,record)
}