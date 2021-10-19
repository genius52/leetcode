package string_issue

func maxVowels(s string, k int) int {
	var aeiou_cnt int = 0
	for i := 0;i < k;i++{
		if s[i] == 'a' ||s[i] == 'e' ||s[i] == 'i'|| s[i] == 'o'|| s[i] == 'u'{
			aeiou_cnt++
		}
	}
	var res int = aeiou_cnt
	if res == k{
		return k
	}
	var l int = len(s)
	var right int = k
	for right < l{
		if s[right - k] == 'a' ||s[right - k] == 'e' ||s[right - k] == 'i'|| s[right - k] == 'o'|| s[right - k] == 'u'{
			aeiou_cnt--
		}
		if s[right] == 'a' ||s[right] == 'e' ||s[right] == 'i'|| s[right] == 'o'|| s[right] == 'u'{
			aeiou_cnt++
		}
		if aeiou_cnt > res{
			res = aeiou_cnt
			if res == k{
				return k
			}
		}
		right++
	}
	return res
}