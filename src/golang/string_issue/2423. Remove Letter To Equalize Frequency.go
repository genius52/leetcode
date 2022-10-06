package string_issue

func equalFrequency(word string) bool {
	var l int = len(word)
	if l == 2{
		return true
	}
	var cnt [26]int
	for i := 0;i < l;i++{
		cnt[word[i] - 'a']++
	}
	var record map[int]int = make(map[int]int)
	for i := 0;i < 26;i++{
		if cnt[i] != 0{
			record[cnt[i]]++
		}
	}
	if len(record) == 1{
		for k,v := range record{
			if k == 1 || v == 1{
				return true
			}else{
				return false
			}
		}
	}
	if len(record) != 2{
		return false
	}
	var min_key int = 1 << 31 - 1
	var max_key int = 0
	for k,_ := range record{
		if k < min_key{
			min_key = k
		}
		if k > max_key{
			max_key = k
		}
	}
	if min_key + 1 == max_key && record[max_key] == 1{
		return true
	}
	if min_key == 1 && record[min_key] == 1{
		return true
	}
	return false
}