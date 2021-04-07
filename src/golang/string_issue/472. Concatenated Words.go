package string_issue

//Input: words = ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]
//Output: ["catsdogcats","dogcatsdog","ratcatdogcat"]
//Explanation: "catsdogcats" can be concatenated by "cats", "dog" and "cats";
//"dogcatsdog" can be concatenated by "dog", "cats" and "dog";
//"ratcatdogcat" can be concatenated by "rat", "cat", "dog" and "cat".
func dfs_findAllConcatenatedWordsInADict(origin string,target string,dict map[string]int,memo map[string]bool)bool{
	var l int = len(target)
	if l == 0{
		return true
	}
	if target != origin{
		if _,ok := dict[target];ok{
			return true
		}
	}
	if _,ok := memo[target];ok{
		return memo[target]
	}
	var res bool = false
	for i := 1;i <= l - 1;i++{
		prefix := target[:i]
		suffix := target[i:]
		if _,ok1 := dict[prefix];ok1{
			if _,ok2 := dict[suffix];ok2{
				res = true
				break
			}else{
				res2 := dfs_findAllConcatenatedWordsInADict(origin,suffix,dict,memo)
				if res2{
					res = true
					break
				}
			}
		}
		if _,ok2 := dict[suffix];ok2{
			res1 := dfs_findAllConcatenatedWordsInADict(origin,prefix,dict,memo)
			if res1{
				res = true
				break
			}
		}
		res1 := dfs_findAllConcatenatedWordsInADict(origin,prefix,dict,memo)
		if !res1{
			continue
		}
		res2 := dfs_findAllConcatenatedWordsInADict(origin,suffix,dict,memo)
		if res1 && res2{
			res = true
			break
		}
	}
	memo[target] = res
	return res
}

func FindAllConcatenatedWordsInADict(words []string) []string {
	//var l int = len(words)
	var dict map[string]int = make(map[string]int)
	for _,w := range words{
		dict[w] = len(w)
	}
	var res []string
	var memo map[string]bool = make(map[string]bool)
	for _,w := range words{
		if len(w) == 0{
			continue
		}
		ret := dfs_findAllConcatenatedWordsInADict(w,w,dict,memo)
		if ret {
			res = append(res,w)
		}
	}
	return res
}