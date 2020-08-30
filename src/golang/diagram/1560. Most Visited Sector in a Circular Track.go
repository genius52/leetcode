package diagram

//Input: n = 4, rounds = [1,3,1,2]
//Output: [1,2]
func MostVisited(n int, rounds []int) []int {
	var l int = len(rounds)
	var record map[int]int = make(map[int]int)
	var res []int
	record[rounds[0]]++
	var max_cnt int = 0
	for i := 1;i < l;i++{
		start := rounds[i - 1]
		end := rounds[i]
		if start > end{
			end += n
		}
		for start < end{
			index := start % n + 1
			record[index]++
			start++
			if record[index] > max_cnt{
				max_cnt = record[index]
			}
		}
	}
	for i := 1;i <= n;i++{
		if cnt,ok := record[i];ok{
			if cnt == max_cnt{
				res = append(res,i)
			}
		}
	}
	return res
}