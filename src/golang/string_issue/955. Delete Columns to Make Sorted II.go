package string_issue

func MinDeletionSize(strs []string) int {
	var rows int = len(strs)
	var columns int = len(strs[0])
	var res int = 0
	var pre_equal []bool = make([]bool,rows)
	for i := 0;i < rows;i++{
		pre_equal[i] = true
	}
	for j := 0;j < columns;j++{
		var cur_equal []bool = make([]bool,rows)
		var valid bool = true
		for i := 1;i < rows;i++{
			if strs[i][j] == strs[i - 1][j]{
				if j == 0{
					cur_equal[i] = true
				}else{
					if pre_equal[i]{
						cur_equal[i] = true
					}else{
						cur_equal[i] = false
					}
				}
			}else if strs[i][j] < strs[i - 1][j]{
				if j == 0{
					res++
					valid = false
					break
				}else{
					if pre_equal[i]{
						res++
						valid = false
						break
					}
				}
			}
		}
		if valid{
			pre_equal = cur_equal
		}
	}
	return res
}