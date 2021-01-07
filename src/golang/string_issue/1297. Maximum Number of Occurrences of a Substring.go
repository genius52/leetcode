package string_issue

//Input: s = "aabcabcab", maxLetters = 2, minSize = 2, maxSize = 3
//Output: 3
func MaxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	var record map[string]int = make(map[string]int)
	var letters map[uint8]int = make(map[uint8]int)
	var l int = len(s)
	var res int = 0
	var begin int = 0
	var end int = 0
	for begin < l{
		var cur_len int = end - begin + 1
		if end == l{
			cur_len = end - begin
		}else{
			if _,ok := letters[s[end]];!ok{
				letters[s[end]] = 1
			}else{
				letters[s[end]]++
			}
		}
		if cur_len >= minSize && cur_len <= maxSize && len(letters) <= maxLetters{
			var sub string
			if end == l{
				sub = s[begin:end]
			}else{
				sub = s[begin:end + 1]
			}
			if _,ok := record[sub];ok{
				record[sub]++
			}else{
				record[sub] = 1
			}
			if record[sub] > res{
				res = record[sub]
			}
			end++
			if end == l{
				begin++
				end = begin
				letters = make(map[uint8]int)
			}
		} else if cur_len < minSize {
			if end == l{
				break
			}
			end++
		}else if cur_len > maxSize || letters[s[end]] > maxLetters{
			begin++
			end = begin
			letters = make(map[uint8]int)
		}
	}
	return res
}