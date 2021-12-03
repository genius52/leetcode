package diagram

//输入：n = 2, languages = [[1],[2],[1,2]], friendships = [[1,2],[1,3],[2,3]]
//输出：1
//解释：你可以选择教用户 1 第二门语言，也可以选择教用户 2 第一门语言。

//输入：n = 3, languages = [[2],[1,3],[1,2],[3]], friendships = [[1,4],[1,2],[3,4],[2,3]]
//输出：2
//解释：教用户 1 和用户 3 第三门语言，需要教 2 名用户。

func MinimumTeachings(n int, languages [][]int, friendships [][]int) int {
	var need_learn_people map[int]bool = make(map[int]bool)
	for _, r := range friendships {
		var find bool = false
		for _, n0 := range languages[r[0]-1] {
		loop:
			for _, n1 := range languages[r[1]-1] {
				if n0 == n1 {
					find = true
					break loop
				}
			}
		}
		if find {
			continue
		}
		need_learn_people[r[0]] = true
		need_learn_people[r[1]] = true
	}
	var res int = 2147483647
	var m int = len(need_learn_people)
	for i := 1; i <= n; i++ {
		var cnt int = 0
		for person, _ := range need_learn_people {
			for _, lan := range languages[person-1] {
				if lan == i {
					cnt++
					break
				}
			}
		}
		res = min_int(res, m-cnt)
	}
	return res
}
