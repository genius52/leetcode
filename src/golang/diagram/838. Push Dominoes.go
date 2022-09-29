package diagram

import "strings"

//Input: ".L.R...LR..L.."
//Output: "LL.RR.LLRRLL.."

//Input: "RR.L"
//Output: "RR.L"

//func PushDominoes(dominoes string) string{
//	var data []string = strings.Split(dominoes,"")
//	data = append([]string{"L"},data...)
//	data = append(data,"R")
//	var l int = len(data)
//	var begin int = 0
//	for begin < l{
//		for begin < l{
//			if data[begin] != "."{
//				break
//			}
//			begin++
//		}
//		var end int = begin + 1
//		for end < l{
//			if data[end] != "."{
//				break
//			}
//			end++
//		}
//		if end >= l{
//			break
//		}
//		if data[begin] == "L" && data[end] == "R"{
//			//do nothing
//
//		}else if data[begin] == "L" && data[end] == "L"{
//			for i := begin + 1;i < end;i++{
//				data[i] = "L"
//			}
//		}else if data[begin] == "R" && data[end] == "R"{
//			for i := begin + 1;i < end;i++{
//				data[i] = "R"
//			}
//		}else if data[begin] == "R" && data[end] == "L"{
//			for i := 1;i <= (end - begin)/2 && (begin + i != end - i);i++{
//				data[begin + i] = "R"
//				data[end - i] = "L"
//			}
//		}
//		begin = end
//	}
//	data = data[1:l - 1]
//	return strings.Join(data,"")
//}

func PushDominoes(dominoes string) string{
	var l int = len(dominoes)
	var right_dis []int = make([]int,l)
	var pre_right int = -1
	for i := 0;i < l;i++{
		if dominoes[i] == 'L'{
			pre_right = -1
		}else if dominoes[i] == 'R'{
			pre_right = i
		}else{
			if pre_right != -1{
				right_dis[i] = i - pre_right
			}
		}
	}
	var left_dis []int = make([]int,l)
	var pre_left int = -1
	for i := l - 1;i >= 0;i--{
		if dominoes[i] == 'L'{
			pre_left = i
		}else if dominoes[i] == 'R'{
			pre_left = -1
		}else{
			if pre_left != -1{
				left_dis[i] = pre_left - i
			}
		}
	}
	var ss strings.Builder
	for i := 0;i < l;i++{
		if dominoes[i] == 'L'{
			ss.WriteRune('L')
			continue
		}
		if dominoes[i] == 'R'{
			ss.WriteRune('R')
			continue
		}
		if left_dis[i] == 0 && right_dis[i] > 0{
			ss.WriteRune('R')
			continue
		}
		if right_dis[i] == 0 && left_dis[i] > 0{
			ss.WriteRune('L')
			continue
		}
		if left_dis[i] > right_dis[i]{
			ss.WriteRune('R')
		}else if left_dis[i] < right_dis[i]{
			ss.WriteRune('L')
		}else{
			ss.WriteRune('.')
		}
	}
	return ss.String()
}