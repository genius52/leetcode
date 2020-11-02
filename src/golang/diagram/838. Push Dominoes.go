package diagram

import "strings"

//Input: ".L.R...LR..L.."
//Output: "LL.RR.LLRRLL.."

//Input: "RR.L"
//Output: "RR.L"

func PushDominoes(dominoes string) string {
	var l int = len(dominoes)
	var data []string = strings.Split(dominoes,"")
	var start int = 0
	var end int = 0
	for start < l && end < l{
		var first_right int = -1
		var last_right int = -1
		for end < l && data[end] != "L"{
			if data[end] == "R"{
				last_right = end
				if first_right == -1{
					first_right = end
				}
			}
			end++
		}
		if end < l{
			if first_right == -1{
				for i := start;i <= end;i++{
					data[i] = "L"
				}
			}else{
				for i := first_right;i <= last_right;i++{
					data[i] = "R"
				}
				for i := 0;last_right + i < end - i;i++{
					data[last_right + i] = "R"
					data[end - i] = "L"
				}
			}
		}else{
			if first_right != -1{
				for i := first_right;i < l;i++{
					data[i] = "R"
				}
			}
		}
		end++
		start = end
	}
	res := strings.Join(data,"")
	return res
}