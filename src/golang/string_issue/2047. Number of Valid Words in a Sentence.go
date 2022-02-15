package string_issue

import "strings"

func countValidWords(sentence string) int {
	var record []string = strings.Split(sentence," ")
	var res int = 0
	var l int = len(record)
	for i := 0;i < l;i++{
		if strings.Trim(record[i]," ") == ""{
			continue
		}
		var match bool = true
		var tag_cnt int = 0
		var tag_idx int = -1
		var connect_cnt int = 0
		var connect_idx int = -1
		for idx,c := range record[i]{
			if c >= '0' && c <= '9'{
				match = false
				break
			}
			if c == '-'{
				connect_cnt++
				connect_idx = idx
				if connect_cnt > 1{
					match = false
					break
				}
			}
			if c == '!' || c == '.' || c == ','{
				tag_idx = idx
				tag_cnt++
				if tag_cnt > 1{
					match = false
					break
				}
			}
		}
		if match{
			if len(record[i]) == 1{
				if connect_idx != -1 {
					continue
				}
			}
			if connect_idx == -1 && tag_idx == -1{
				res++
			}else if connect_idx == -1 && tag_idx == len(record[i]) - 1{
				res++
			}else if tag_idx == -1 && connect_idx != 0 && connect_idx != len(record[i]) - 1{
				var left int = connect_idx - 1
				var right int = connect_idx + 1
				var find_left bool = false
				for left >= 0{
					if record[i][left] >= 'a' && record[i][left] <= 'z'{
						find_left = true
						break
					}
					left--
				}
				var find_right bool = false
				for right < len(record[i]){
					if record[i][right] >= 'a' && record[i][right] <= 'z'{
						find_right = true
						break
					}
					right++
				}
				if find_left && find_right{
					res++
				}
			}else if connect_idx != 0 && connect_idx != len(record[i]) - 1 && tag_idx == len(record[i]) - 1{
				var left int = connect_idx - 1
				var right int = connect_idx + 1
				var find_left bool = false
				for left >= 0{
					if record[i][left] >= 'a' && record[i][left] <= 'z'{
						find_left = true
						break
					}
					left--
				}
				var find_right bool = false
				for right < len(record[i]){
					if record[i][right] >= 'a' && record[i][right] <= 'z'{
						find_right = true
						break
					}
					right++
				}
				if find_left && find_right{
					res++
				}
			}
		}
	}
	return res
}