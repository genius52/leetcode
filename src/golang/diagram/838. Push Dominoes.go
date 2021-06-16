package diagram

import "strings"

//Input: ".L.R...LR..L.."
//Output: "LL.RR.LLRRLL.."

//Input: "RR.L"
//Output: "RR.L"

func PushDominoes(dominoes string) string{
	var data []string = strings.Split(dominoes,"")
	data = append([]string{"L"},data...)
	data = append(data,"R")
	var l int = len(data)
	var begin int = 0
	for begin < l{
		for begin < l{
			if data[begin] != "."{
				break
			}
			begin++
		}
		var end int = begin + 1
		for end < l{
			if data[end] != "."{
				break
			}
			end++
		}
		if end >= l{
			break
		}
		if data[begin] == "L" && data[end] == "R"{
			//do nothing

		}else if data[begin] == "L" && data[end] == "L"{
			for i := begin + 1;i < end;i++{
				data[i] = "L"
			}
		}else if data[begin] == "R" && data[end] == "R"{
			for i := begin + 1;i < end;i++{
				data[i] = "R"
			}
		}else if data[begin] == "R" && data[end] == "L"{
			for i := 1;i <= (end - begin)/2 && (begin + i != end - i);i++{
				data[begin + i] = "R"
				data[end - i] = "L"
			}
		}
		begin = end
	}
	data = data[1:l - 1]
	return strings.Join(data,"")
}

//func PushDominoes(dominoes string) string {
//	var l int = len(dominoes)
//	var data []string = strings.Split(dominoes,"")
//	var start int = 0
//	var end int = 0
//	for start < l && end < l{
//		var first_right int = -1
//		var last_right int = -1
//		for end < l && data[end] != "L"{
//			if data[end] == "R"{
//				last_right = end
//				if first_right == -1{
//					first_right = end
//				}
//			}
//			end++
//		}
//		if end < l{
//			if first_right == -1{
//				for i := start;i <= end;i++{
//					data[i] = "L"
//				}
//			}else{
//				for i := first_right;i <= last_right;i++{
//					data[i] = "R"
//				}
//				for i := 0;last_right + i < end - i;i++{
//					data[last_right + i] = "R"
//					data[end - i] = "L"
//				}
//			}
//		}else{
//			if first_right != -1{
//				for i := first_right;i < l;i++{
//					data[i] = "R"
//				}
//			}
//		}
//		end++
//		start = end
//	}
//	res := strings.Join(data,"")
//	return res
//}