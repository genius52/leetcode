package string_issue

type WordFilter struct {
	prefix map[string][]int
	suffix map[string][]int
}

func Constructor745(words []string) WordFilter {
	var obj WordFilter
	obj.prefix = make(map[string][]int)
	obj.suffix = make(map[string][]int)
	var l int = len(words)
	for i := l - 1;i >= 0;i--{
		var l2 int = len(words[i])
		for j := 0;j < l2;j++{
			if _,ok := obj.prefix[words[i][:j + 1]];ok{
				obj.prefix[words[i][:j + 1]] = append(obj.prefix[words[i][:j + 1]],i)
			}else{
				obj.prefix[words[i][:j + 1]] = []int{i}
			}
			if _,ok := obj.suffix[words[i][j:]];ok{
				obj.suffix[words[i][j:]] = append(obj.suffix[words[i][j:]],i)
			}else{
				obj.suffix[words[i][j:]] = []int{i}
			}
		}
	}
	return obj
}

func (this *WordFilter) F(prefix string, suffix string) int {
	if _,ok := this.prefix[prefix];!ok{
		return -1
	}
	if _,ok := this.suffix[suffix];!ok{
		return -1
	}
	for _,idx1 := range this.prefix[prefix]{
		for _,idx2 := range this.suffix[suffix]{
			if idx1 == idx2{
				return idx1
			}
		}
	}
	return -1
}