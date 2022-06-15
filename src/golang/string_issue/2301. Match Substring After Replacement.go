package string_issue

//func dfs_matchReplacement(s string,sub string,l int,idx int,pre string,old_new map[uint8][]uint8)bool{
//	if idx >= l{
//		return true
//	}
//	cur := pre + string(sub[idx])
//	if strings.Contains(s,cur){
//		ret := dfs_matchReplacement(s,sub,l,idx + 1,cur,old_new)
//		if ret{
//			return true
//		}
//	}
//	if vals,ok := old_new[sub[idx]];ok{
//		for _,c := range vals{
//			cur2 := pre + string(c)
//			if strings.Contains(s,cur2){
//				ret := dfs_matchReplacement(s,sub,l,idx + 1,cur2,old_new)
//				if ret{
//					return true
//				}
//			}
//		}
//	}
//	return false
//}
//
//func MatchReplacement(s string, sub string, mappings [][]byte) bool {
//	var old_new map[uint8][]uint8 = make(map[uint8][]uint8)
//	for _,m := range mappings{
//		old_new[m[0]] = append(old_new[m[0]],m[1])
//	}
//	return dfs_matchReplacement(s,sub,len(sub),0,"",old_new)
//}

func MatchReplacement(s string, sub string, mappings [][]byte) bool {
	var old_new map[uint8][]uint8 = make(map[uint8][]uint8)
	for _,m := range mappings{
		old_new[m[0]] = append(old_new[m[0]],m[1])
	}
	var l1 int = len(s)
	var l2 int = len(sub)
	for i := 0;i <= l1 - l2;i++{
		j := 0
		for ;j < l2;j++{
			if s[i + j] == sub[j]{
				continue
			}
			if _,ok := old_new[sub[j]];!ok{
				break
			}
			var find bool = false
			for _,c := range old_new[sub[j]]{
				if s[i + j] == c{
					find = true
					break
				}
			}
			if !find{
				break
			}
		}
		if j == l2{
			return true
		}
	}
	return false
}