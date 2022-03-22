package string_issue

func max_int64(a int64,b int64)int64{
	if a > b{
		return a
	}
	return b
}

func MaximumSubsequenceCount(text string, pattern string) int64 {
	var l int = len(text)
	var prefix_first []int64 = make([]int64,l)
	var suffix_second []int64 = make([]int64,l)
	if text[0] == pattern[0]{
		prefix_first[0] = 1
	}
	for i := 1;i < l;i++{
		if text[i] == pattern[0]{
			prefix_first[i] = 1 + prefix_first[i - 1]
		}else{
			prefix_first[i] = prefix_first[i - 1]
		}
	}
	if text[l - 1] == pattern[1]{
		suffix_second[l - 1] = 1
	}
	for i := l - 2;i >= 0;i--{
		if text[i] == pattern[1]{
			suffix_second[i] = 1 + suffix_second[i + 1]
		}else{
			suffix_second[i] = suffix_second[i + 1]
		}
	}
	var res int64 = 0
	for i := 0;i < l - 1;i++{
		if text[i] == pattern[0]{
			res += suffix_second[i + 1]
		}
	}
	var add int64 = 0
	if l == 1{
		add = max_int64(suffix_second[0],prefix_first[0])
	}else{
		for i := 0;i < l;i++{
			var cur int64 = 0
			if i == 0{
				cur = 1 * suffix_second[0]
			}else if i == l - 1{
				cur = prefix_first[i] * 1
			}else{
				cur = max_int64(suffix_second[i + 1],prefix_first[i])
			}
			add = max_int64(add,cur)
		}
	}
	return res + add
}