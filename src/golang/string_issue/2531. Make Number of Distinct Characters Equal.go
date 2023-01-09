package string_issue

func IsItPossible(word1 string, word2 string) bool {
	var l1 int = len(word1)
	var l2 int = len(word2)
	var cnt1 [26]int
	for i := 0;i < l1;i++{
		cnt1[word1[i] - 'a']++
	}
	var cnt2 [26]int
	for i := 0;i < l2;i++{
		cnt2[word2[i] - 'a']++
	}
	var kind1 int = 0
	var kind2 int = 0
	for i := 0;i < 26;i++{
		if cnt1[i] > 0{
			kind1++
		}
		if cnt2[i] > 0{
			kind2++
		}
	}
	for i := 0;i < 26;i++{
		if cnt1[i] == 0{
			continue
		}
		for j := 0;j < 26;j++{
			if cnt2[j] == 0{
				continue
			}
			if i == j{
				if kind1 == kind2{
					return true
				}else{
					continue
				}
			}
			var new_kind1 int = kind1
			var new_kind2 int = kind2
			if cnt1[j] == 0{
				new_kind1++
			}
			if cnt1[i] == 1{
				new_kind1--
			}
			if cnt2[i] == 0{
				new_kind2++
			}
			if cnt2[j] == 1{
				new_kind2--
			}
			if new_kind1 == new_kind2{
				return true
			}
		}
	}
	return false
}