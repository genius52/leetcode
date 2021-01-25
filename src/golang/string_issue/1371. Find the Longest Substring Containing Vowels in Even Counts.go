package string_issue

//Input: s = "eleetminicoworoep"
//Output: 13
//Explanation: The longest substring is "leetminicowor" which contains two each of the vowels: e, i and o and zero of the vowels: a and u.

func FindTheLongestSubstring(s string) int {
	var a_cnt int = 0
	var e_cnt int = 0
	var i_cnt int = 0
	var o_cnt int = 0
	var u_cnt int = 0
	var l int = len(s)
	var record map[int]int = make(map[int]int)
	record[0] = -1
	var max_len int = 0
	for i := 0;i < l;i++{
		if s[i] == 'a'{
			a_cnt++
		}else if s[i] == 'e'{
			e_cnt++
		}else if s[i] == 'i'{
			i_cnt++
		}else if s[i] == 'o'{
			o_cnt++
		}else if s[i] == 'u'{
			u_cnt++
		}

		var cur_key int = 0
		if (a_cnt | 1) == a_cnt{
			cur_key |= (a_cnt % 2) << 4
		}
		if (e_cnt | 1) == e_cnt{
			cur_key |= (e_cnt % 2) << 3
		}
		if (i_cnt | 1) == i_cnt{
			cur_key |= (i_cnt % 2) << 2
		}
		if (o_cnt | 1) == o_cnt{
			cur_key |= (o_cnt % 2) << 1
		}
		if (u_cnt | 1) == u_cnt{
			cur_key |= (u_cnt % 2)
		}
		if pos,ok := record[cur_key];ok{
			cur_len := i - pos
			if cur_len > max_len{
				max_len = cur_len
			}
		}else{
			record[cur_key] = i
		}
	}
	return max_len
}