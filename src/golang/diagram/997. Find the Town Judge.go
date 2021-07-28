package diagram


//Input: N = 4, trust = [[1,3],[1,4],[2,3],[2,4],[4,3]]
//Output: 3
//3
//[[[1,3],[2,3],[3,1]]
func findJudge2(N int, trust [][]int) int{
	var scores []int = make([]int,N + 1)
	for _,t := range trust{
		scores[t[1]]++
		scores[t[0]]--
	}
	for i := 1;i <= N;i++{
		if scores[i] == (N - 1){
			return i
		}
	}
	return -1
}

func findJudge(N int, trust [][]int) int {
	var len int = len(trust)
	var record []int = make([]int,N+1)
	var i int = 0
	for i < len{
		record[trust[i][0]] = -1
		if record[trust[i][1]] != -1{
			record[trust[i][1]]++
		}
		i++
	}
	for index,e := range record{
		if e == (N-1){
			return index
		}
	}
	return -1
}