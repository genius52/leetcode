package string_issue

import "strconv"

func CountPalindromes(s string) int {
	var l int = len(s)
	var prefix_onechar [][10]int = make([][10]int,l)
	var prefix_twochar []map[string]int = make([]map[string]int,l)
	var suffix_onechar [][10]int = make([][10]int,l)
	var suffix_twochar []map[string]int = make([]map[string]int,l)
	var MOD int = 1e9 + 7
	for i := 0;i < l;i++{
		prefix_twochar[i] = make(map[string]int)
		if i == 0{
			prefix_onechar[i][int(s[i] - '0')]++
			continue
		}
		for k,v := range prefix_twochar[i - 1]{
			prefix_twochar[i][k] = v
		}
		for j := 0;j <= 9;j++{
			prefix_onechar[i][j] = prefix_onechar[i - 1][j]
			if prefix_onechar[i - 1][j] > 0{
				prefix_twochar[i][strconv.Itoa(j) + string(s[i])] += prefix_onechar[i - 1][j]
				prefix_twochar[i][strconv.Itoa(j) + string(s[i])] %= MOD
			}
		}
		prefix_onechar[i][int(s[i] - '0')]++
	}
	for i := l - 1;i >= 0;i--{
		suffix_twochar[i] = make(map[string]int)
		if i == l - 1{
			suffix_onechar[i][int(s[i] - '0')]++
			continue
		}
		for k,v := range suffix_twochar[i + 1]{
			suffix_twochar[i][k] = v
		}
		for j := 0;j <= 9;j++{
			suffix_onechar[i][j] = suffix_onechar[i + 1][j]
			if suffix_onechar[i + 1][j] > 0{
				suffix_twochar[i][strconv.Itoa(j) + string(s[i])] += suffix_onechar[i + 1][j]
				suffix_twochar[i][strconv.Itoa(j) + string(s[i])] %= MOD
			}
		}
		suffix_onechar[i][int(s[i] - '0')]++
	}
	var record map[string]int = make(map[string]int)

	for i := 2;i < l - 2;i++{
		for n1 := 0;n1 <= 9;n1++{
			for n2 := 0;n2 <= 9;n2++{
				key := strconv.Itoa(n1) + strconv.Itoa(n2)
				if val1,ok1 := prefix_twochar[i - 1][key];ok1{
					if val2,ok2 := suffix_twochar[i + 1][key];ok2{
						record[key + string(s[i]) + key] += val1 * val2
						record[key + string(s[i]) + key] %= MOD
					}
				}
			}
		}
	}
	var res int = 0
	for _,v := range record{
		res += v
		res %= MOD
	}
	return res
}