package string_issue

type Tie struct{
	data [26]*Tie
	//group map[int]int //
	cnt int
}

func sumPrefixScores(words []string) []int {
	var l int = len(words)
	var root Tie
	//root.group = make(map[int]int)
	for i := 0;i < l;i++{
		var visit *Tie = &root
		for j := 0;j < len(words[i]);j++{
			if visit.data[words[i][j] - 'a'] == nil{
				visit.data[words[i][j] - 'a'] = new(Tie)
				//visit.data[words[i][j] - 'a'].group = make(map[int]int)
			}
			//visit.data[words[i][j] - 'a'].group[i]++
			visit.data[words[i][j] - 'a'].cnt++
			visit = visit.data[words[i][j] - 'a']
		}
	}
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		var cnt int = 0
		var visit *Tie = &root
		for j := 0;j < len(words[i]);j++{
			cnt += visit.data[words[i][j] - 'a'].cnt
			visit = visit.data[words[i][j] - 'a']
		}
		res[i] = cnt
	}
	return res
}