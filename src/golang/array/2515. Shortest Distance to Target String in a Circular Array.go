package array

func closetTarget(words []string, target string, startIndex int) int {
	var word_idx map[string][]int = make(map[string][]int)
	for i,w := range words{
		word_idx[w] = append(word_idx[w],i)
	}
	if _,ok := word_idx[target];!ok{
		return -1
	}
	var l int = len(words)
	var res int = l
	for _,idx := range word_idx[target]{
		dis := min_int(min_int(idx,startIndex) + l - max_int(idx,startIndex),abs_int(idx - startIndex))
		if dis < res{
			res = dis
		}
	}
	return res
}