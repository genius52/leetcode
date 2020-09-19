package string_issue

type MagicDictionary struct {
	origin map[string]bool
	transform map[string]int
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	var obj MagicDictionary
	obj.origin = make(map[string]bool)
	obj.transform = make(map[string]int)
	return obj
}

func (this *MagicDictionary) BuildDict(dictionary []string)  {
	for _,s := range dictionary{
		this.origin[s] = true
		var l int = len(s)
		for i := 0;i < l;i++{
			var tmp string
			for j := 0;j < l;j++{
				if j == i{
					tmp += "*"
				}else{
					tmp += string(s[j])
				}
			}
			this.transform[tmp]++
		}
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	//var find_origin bool = false
	//if _,ok := this.origin[searchWord];ok{
	//	find_origin = true
	//}
	var l int = len(searchWord)
	for i := 0;i < l;i++{
		var tmp string
		for j := 0;j < l;j++{
			if j == i{
				tmp += "*"
			}else{
				tmp += string(searchWord[j])
			}
		}
		if cnt,ok := this.transform[tmp];ok{
			if cnt >= 2{
				return true
			}
			if _,ok2 := this.origin[searchWord];!ok2{
				return true
			}
		}
	}
	return false
}

