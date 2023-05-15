package number

func circularGameLosers(n int, k int) []int {
	var visited []bool = make([]bool, n)
	var idx int = 0
	var steps int = 1
	visited[idx] = true
	for true {
		next := (idx + steps*k) % n
		if visited[next] {
			break
		}
		visited[next] = true
		idx = next
		steps++
	}
	var res []int
	for i := 0; i < n; i++ {
		if !visited[i] {
			res = append(res, i+1)
		}
	}
	return res
}
