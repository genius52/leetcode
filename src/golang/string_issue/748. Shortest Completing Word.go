package string_issue

func shortestCompletingWord(licensePlate string, words []string) string {
	var need [26]int
	for _,c := range licensePlate{
		if c >= 'a' && c <= 'z'{
			need[c - 'a']++
		}else if c >= 'A' && c <= 'Z'{
			need[c - 'A']++
		}
	}
	var res string
	var pre_length int = 2147483647
	for _,w := range words{
		var copy [26]int = need
		for _,c := range w{
			copy[c - 'a']--
		}
		var complete bool = true
		for i := 0;i < 26;i++{
			if copy[i] > 0{
				complete = false
				break
			}
		}
		if complete{
			cur_len := len(w)
			if cur_len < pre_length{
				pre_length = cur_len
				res = w
			}
		}
	}
	return res
}