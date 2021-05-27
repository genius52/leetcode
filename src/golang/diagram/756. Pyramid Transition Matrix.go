package diagram

//输入：bottom = "BCD", allowed = ["BCG", "CDE", "GEA", "FFF"]
//输出：true
//解释：
//可以堆砌成这样的金字塔:
//    A
//   / \
//  G   E
// / \ / \
//B   C   D
//
//因为符合 BCG、CDE 和 GEA 三种规则
func dfs_pyramidTransition(s string,dict map[string][]uint8)bool{
	var l int = len(s)
	if l == 1{
		return true
	}
	var pre []string
	for i := 0;i < l - 1;i++{
		sub := s[i:i + 2]
		if chars,ok := dict[sub];!ok{
			return false
		}else{
			var cur []string
			for _,c := range chars{
				if len(pre) == 0{
					cur = append(cur,string(c))
				}else{
					for _,p := range pre{
						cur = append(cur,p + string(c))
					}
				}

			}
			pre = cur
		}
	}
	for _,n := range pre{
		r := dfs_pyramidTransition(n,dict)
		if r{
			return true
		}
	}
	return false
}

func PyramidTransition(bottom string, allowed []string) bool {
	var dict map[string][]uint8 = make(map[string][]uint8)
	for _,word := range allowed{
		if _,ok := dict[word[:2]];ok{
			dict[word[:2]] = append(dict[word[:2]],word[2])
		}else{
			dict[word[:2]] = []uint8{word[2]}
		}
	}
	return dfs_pyramidTransition(bottom,dict)
}