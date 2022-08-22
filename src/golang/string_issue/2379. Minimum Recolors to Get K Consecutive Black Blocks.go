package string_issue

func MinimumRecolors(blocks string, k int) int {
	var l int = len(blocks)
	var white_cnt int = 0
	for i := 0;i < k;i++{
		if blocks[i] == 'W'{
			white_cnt++
		}
	}
	var res int = white_cnt
	var i int = 0
	for i + k < l{
		if blocks[i] == 'W'{
			white_cnt--
		}
		if blocks[i + k] == 'W'{
			white_cnt++
		}
		if white_cnt < res{
			res = white_cnt
		}
		i++
	}
	return res
}