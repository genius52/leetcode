package string_issue

//Input: stickers = ["with","example","science"], target = "thehat"
//Output: 3
//Explanation:
//We can use 2 "with" stickers, and 1 "example" sticker.
//After cutting and rearrange the letters of those stickers, we can form the target "thehat".
//Also, this is the minimum number of stickers necessary to form the target string.

func dp_minStickers(record [][26]int,l int,target string,memo map[string]int)int{
	if len(target) == 0{
		return 0
	}
	if _,ok := memo[target];ok{
		return memo[target]
	}
	var res int = 2147483647
	memo[target] = res
	for i := 0;i < l;i++{
		var s string
		var used [26]int = record[i]
		for k := 0;k < len(target);k++{
			if used[target[k] - 'a'] > 0{
				used[target[k] - 'a']--
			}else{
				s += string(target[k])
			}
		}
		next := dp_minStickers(record,l,s,memo)
		if next != 2147483647{
			res = min_int(res,1 + next)
		}
	}
	memo[target] = res
	return res
}

func MinStickers(stickers []string, target string) int {
	var l int = len(stickers)
	var record [][26]int = make([][26]int,l)
	for i := 0;i < l;i++{
		for j := 0;j < len(stickers[i]);j++{
			record[i][stickers[i][j] - 'a']++
		}
	}
	var memo map[string]int = make(map[string]int)
	res := dp_minStickers(record,l,target,memo)
	if res == 2147483647{
		return -1
	}
	return res
}

//func dp_minStickers(stickers []string,l int,pos int,cnt [26]int)int{
//	var empty bool = true
//	var max_occur_cnt int = 0
//	for i := 0;i < 26;i++{
//		if cnt[i] > 0{
//			empty = false
//		}
//		if cnt[i] > max_occur_cnt{
//			max_occur_cnt = cnt[i]
//		}
//	}
//	if empty{
//		return 0
//	}
//	if pos >= l{
//		return -1
//	}
//
//	var base_cnt [26]int
//	for _,c := range stickers[pos]{
//		base_cnt[c - 'a']++
//	}
//	var min_steps int = -1
//	var pre [26]int
//	for j := 1;j <= max_occur_cnt;j++ {
//		var multiple [26]int
//		for i := 0;i < 26;i++{
//			multiple[i] = base_cnt[i] * j
//		}
//		var copy_cnt [26]int = cnt
//		var used bool = false
//		for i := 0;i < 26;i++{
//			if copy_cnt[i] > 0 && multiple[i] > 0{
//				used = true
//			}
//			if multiple[i] >= copy_cnt[i]{
//				copy_cnt[i] = 0
//			}else{
//				copy_cnt[i] -= multiple[i]
//			}
//		}
//		if !used{
//			break
//		}
//		if copy_cnt == pre{
//			break
//		}
//		pre = copy_cnt
//		next_steps := dp_minStickers(stickers,l,pos + 1,copy_cnt)
//		if next_steps == -1{
//			continue
//		}
//		if min_steps == -1 || (next_steps + j) < min_steps{
//			min_steps = next_steps + j
//		}
//	}
//	//skip cur word
//	skip_cur_steps := dp_minStickers(stickers,l,pos + 1,cnt)
//	//choose cur word
//	if skip_cur_steps == -1{
//		return min_steps
//	}
//	if min_steps == -1{
//		return skip_cur_steps
//	}
//	return min_int(min_steps,skip_cur_steps)
//}
//
//func MinStickers(stickers []string, target string) int {
//	var cnt [26]int
//	for _,c := range target{
//		cnt[c - 'a']++
//	}
//	return dp_minStickers(stickers,len(stickers),0,cnt)
//}