package array

import "strconv"

func PrisonAfterNDays(cells []int, n int) []int {
	var visited map[string]bool = make(map[string]bool)
	var k int = n
	for k > 0{
		var cur []int = make([]int,8)
		for i := 1;i < 7;i++{
			if cells[i - 1] == cells[i + 1]{
				cur[i] = 1
			}
		}
		var s2 string
		for i := 0;i < 8;i++{
			s2 += strconv.Itoa(cur[i])
		}
		if _,ok := visited[s2];ok{
			break
		}
		cells = cur
		visited[s2] = true
		k--
	}
	if k == 0{
		return cells
	}
	n = n % (n - k)
	for n > 0{
		var cur []int = make([]int,8)
		for i := 1;i < 7;i++{
			if cells[i - 1] == cells[i + 1]{
				cur[i] = 1
			}
		}
		cells = cur
		n--
	}
	return cells
}