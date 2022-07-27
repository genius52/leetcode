package array

func shortestSequence(rolls []int, k int) int {
	var l int = len(rolls)
	var visited map[int]bool = make(map[int]bool)
	var res int = 0
	for i := 0;i < l;i++{
		visited[rolls[i]] = true
		if len(visited) == k{
			res++
			visited = make(map[int]bool)
		}
	}
	return res
}