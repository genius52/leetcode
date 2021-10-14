package number

import "strconv"

func MaxDiff(num int) int {
	s := strconv.Itoa(num)
	var l int = len(s)
	var max_s string
	var find bool = false
	var find_num uint8
	for i := 0;i < l;i++{
		if find{
			if s[i] == find_num{
				max_s += "9"
			}else{
				max_s += string(s[i])
			}
		}else{
			max_s += "9"
			if s[i] != '9'{
				find_num = s[i]
				find = true
			}
		}
	}
	var min_s string = "1"
	find = false
	//var replace string
	var match uint8
	//if s[0] != '1'{
	//	replace = "1"
	//}
	for i := 1;i < l;i++{
		if s[0] == '1'{//'1' will never replace
			if match != 0{
				if s[i] == match{
					min_s += "0"
				}else{
					min_s += string(s[i])
				}
			}else{
				if s[i] != '1' && s[i] != '0'{
					//replace = string(s[i])
					match = s[i]
					min_s += "0"
				}else{
					min_s += string(s[i])
				}
			}
		}else{
			if s[i] == s[0]{
				min_s += "1"
			}else{
				min_s += string(s[i])
			}
		}
	}
	big,_ := strconv.Atoi(max_s)
	small,_ := strconv.Atoi(min_s)
	return big - small
}