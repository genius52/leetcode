package string_issue

//Input: text = "aaabaaa"
//Output: 6
func MaxRepOpt1(text string) int {
	var record map[uint8]int = make(map[uint8]int)
	var l int = len(text)
	for i := 0;i < l;i++{
		record[text[i]]++
	}
	var chars []uint8
	var counts []int
	var begin int = 0
	var end int = 0
	var max_len int = 0
	for end < l{
		if text[end] != text[begin]{
			chars = append(chars,text[begin])
			cur_len := end - begin
			counts = append(counts,cur_len)
			begin = end
			if cur_len > max_len{
				max_len = cur_len
			}
		}else{
			end++
		}
	}
	chars = append(chars,text[l - 1])
	counts = append(counts,end - begin)
	if end - begin > max_len{
		max_len = end - begin
	}
	var l2 int = len(chars)
	//case 1
	for i := 0;i < l2;i++{
		if record[chars[i]] - counts[i] > 0{
			cur_len := counts[i] + 1
			if cur_len > max_len{
				max_len = cur_len
			}
		}
	}
	//case 2
	for i := 1;i < l2 - 1;i++{
		if chars[i - 1] != chars[i + 1] || counts[i] != 1{
			continue
		}
		rest_len := record[chars[i - 1]] - counts[i - 1] - counts[i + 1]
		var cur_len int = 0
		if rest_len >= 1{
			cur_len =  counts[i - 1] + counts[i + 1] + 1
		}else{
			cur_len =  counts[i - 1] + counts[i + 1]
		}
		if cur_len > max_len{
			max_len = cur_len
		}
	}
	return max_len
}