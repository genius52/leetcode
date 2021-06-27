package string_issue

func WonderfulSubstrings(word string) int64 {
	var l int = len(word)
	var record map[int]int64 = make(map[int]int64)
	var res int64 = 0
	var prefix []int = make([]int,l + 1)
	record[0] = 1
	for i := 0;i < l;i++{
		if i == 0{
			prefix[i] = 1 << int(word[i] - 'a')
		}else{
			prefix[i] = prefix[i - 1] ^ (1 << int(word[i] - 'a'))
		}
		if cnt,ok := record[prefix[i]];ok{
			res += cnt
		}
		for j := 0;j < 10;j++{
			mask := 1 << j
			if cnt,ok := record[prefix[i] ^ mask];ok{
				res += cnt
			}
		}
		//TLE
		//for n,cnt := range record{
		//	n2 := n ^ prefix[i]
		//	if n2 == 0 || (n2 & (n2 - 1)) == 0{
		//		res += cnt
		//	}
		//}
		if _,ok := record[prefix[i]];ok{
			record[prefix[i]]++
		}else{
			record[prefix[i]] = 1
		}
	}
	return res
}