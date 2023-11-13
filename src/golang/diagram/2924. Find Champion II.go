package diagram

func findChampion(n int, edges [][]int) int {
	var indgree []int = make([]int, n)
	for _, edge := range edges {
		indgree[edge[1]]++
	}
	var team int = -1
	for i := 0; i < n; i++ {
		if indgree[i] == 0 {
			if team != -1 {
				return -1
			}
			team = i
		}
	}
	return team
}
