package number

func numberOfWays(s string) int64 {
	var l int = len(s)
	var prefix_zero []int64 = make([]int64,l + 1)
	var prefix_one []int64 = make([]int64,l + 1)
	var suffix_zero []int64 = make([]int64,l + 1)
	var suffix_one []int64 = make([]int64,l + 1)
	for i := 0;i < l;i++{
		if s[i] == '0'{
			prefix_zero[i + 1] = 1 + prefix_zero[i]
			prefix_one[i + 1] = prefix_one[i]
		}else if s[i] == '1'{
			prefix_one[i + 1] = 1 + prefix_one[i]
			prefix_zero[i + 1] = prefix_zero[i]
		}
	}
	for i := l - 1;i >= 0;i--{
		if s[i] == '0'{
			suffix_zero[i] = 1 + suffix_zero[i + 1]
			suffix_one[i] = suffix_one[i + 1]
		}else if s[i] == '1'{
			suffix_one[i] = 1 + suffix_one[i + 1]
			suffix_zero[i] = suffix_zero[i + 1]
		}
	}
	var res int64 = 0
	for i := 1;i < l - 1;i++{
		if s[i] == '0'{
			res += prefix_one[i] * suffix_one[i]
		}else if s[i] == '1'{
			res += prefix_zero[i] * suffix_zero[i]
		}
	}
	return res
}