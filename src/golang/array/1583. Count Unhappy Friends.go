package array

//Input: n = 4, preferences = [[1, 2, 3], [3, 2, 0], [3, 1, 0], [1, 2, 0]], pairs = [[0, 1], [2, 3]]
//Output: 2
//Explanation:
//Friend 1 is unhappy because:
//- 1 is paired with 0 but prefers 3 over 0, and
//- 3 prefers 1 over 2.
//Friend 3 is unhappy because:
//- 3 is paired with 2 but prefers 1 over 2, and
//- 1 prefers 3 over 0.
//Friends 0 and 2 are happy.
func UnhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	var relation [][]int = make([][]int,n)
	for i := 0;i < n;i++{
		relation[i] = make([]int,n)
	}
	for i := 0;i < n;i++{
		for j := 0;j < n - 1;j++{
			relation[i][preferences[i][j]] = j
		}
	}
	var group map[int]int = make(map[int]int)
	for _,p := range pairs{
		group[p[0]] = p[1]
		group[p[1]] = p[0]
	}
	var res int = 0
	for _,p := range pairs{
		if relation[p[0]][p[1]] > 0{
			for i := 0;i < n;i++{
				if i == p[0] || i == p[1] || relation[p[0]][i] > relation[p[0]][p[1]]{
					continue
				}
				if relation[i][p[0]] < relation[i][group[i]]{
					res++
					break
				}
			}
		}
		if relation[p[1]][p[0]] > 0{
			for i := 0;i < n;i++{
				if i == p[0] || i == p[1] || relation[p[1]][i] > relation[p[1]][p[0]]{
					continue
				}
				if relation[i][p[1]] < relation[i][group[i]]{
					res++
					break
				}
			}
		}
	}
	return res
}