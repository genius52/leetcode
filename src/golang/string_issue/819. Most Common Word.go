package string_issue

import "sort"

func MostCommonWord(paragraph string, banned []string) string {
	var forbidden map[string]bool = make(map[string]bool)
	for _,word := range banned{
		forbidden[word] = true
	}
	var word_cnt map[string]int = make(map[string]int)
	var l int = len(paragraph)
	var left int = 0
	for left < l{
		var right int = left
		for right < l && ((paragraph[right] >= 'a' && paragraph[right] <= 'z') ||
			(paragraph[right] >= 'A' && paragraph[right] <= 'Z')){
			right++
		}
		cur := toLowerCase(paragraph[left:right])
		if _,ok := word_cnt[cur];ok{
			word_cnt[cur]++
		}else{
			word_cnt[cur] = 1
		}
		left = right + 1
		for left < l{
			if (paragraph[left] >= 'a' && paragraph[left] <= 'z') ||
				(paragraph[left] >= 'A' && paragraph[left] <= 'Z'){
				break
			}
			left++
		}
	}
	var cnt_word map[int][]string = make(map[int][]string)
	for word,cnt := range word_cnt{
		cnt_word[cnt] = append(cnt_word[cnt],word)
	}
	var keys []int
	for cnt,_ := range cnt_word{
		keys = append(keys,cnt)
	}
	sort.Ints(keys)
	var key_l int = len(keys)
	for i := key_l - 1;i >= 0;i--{
		var frequency int = keys[i]
		for _,word := range cnt_word[frequency]{
			if _,ok := forbidden[word];!ok{
				return toLowerCase(word)
			}
		}
	}
	return ""
}