package string_issue

func ReorganizeString(S string) string {
	var l int = len(S)
	var record map[uint8]int = make(map[uint8]int)
	var max_cnt int = 0
	var letter uint8 = 0
	for i := 0;i < l;i++{
		record[S[i]]++
		if record[S[i]] > max_cnt{
			letter = S[i]
			max_cnt = record[S[i]]
		}
	}
	if max_cnt > (l + 1) /2{
		return ""
	}
	var data []uint8 = make([]uint8,l)
	var index int = 0
	for i := 0;i < record[letter];i++{
		data[index] = letter
		index += 2
	}
	delete(record,letter)

	for c,cnt := range record{
		for i := 0;i < cnt;i++{
			if index >= l{
				index = 1
			}
			data[index] = c
			index += 2
		}
	}
	var res string
	for _,c := range data{
		res += string(c)
	}
	return res
}