package string_issue

func DecodeAtIndex(S string, K int) string{
	var diff int = 0
	for i := 0;diff < K;i++{
		if S[i] >= '0' && S[i] <= '9'{
			diff = diff * int((S[i] - uint8('0')))
		}else{
			diff++
		}
	}
	return ""
}

//func DecodeAtIndex(S string, K int) string {
//	var data []uint8 = make([]uint8,K)
//	var index int = 0
//	var l int = len(S)
//	for i := 0;i < l;i++{
//		if S[i] >= '0' && S[i] <= '9'{
//			cnt,_ := strconv.Atoi(string(S[i]))
//			var end int = index - 1
//			for i := 0;i < cnt - 1;i++{
//				for j := 0;j <= end;j++{
//					data[index] = data[j]
//					if index == K - 1 {
//						return string(data[index])
//					}
//					index++
//				}
//			}
//		}else{
//			if index == K{
//				break
//			}
//			data[index] = S[i]
//			index++
//			if index == K{
//				break
//			}
//		}
//	}
//	return string(data[K - 1])
//}