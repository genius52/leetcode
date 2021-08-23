package array

func numEquivDominoPairs(dominoes [][]int) int{
	var l int = len(dominoes)
	var record map[int]int = make(map[int]int)
	var res int = 0
	for i := 0;i < l;i++{
		key := min_int(dominoes[i][0],dominoes[i][1]) * 10 + max_int(dominoes[i][0],dominoes[i][1])
		if _,ok := record[key];ok{
			record[key]++
		}else{
			record[key] = 1
		}
	}
	for _,cnt := range record{
		res += cnt * (cnt - 1)/2
	}
	return res
}