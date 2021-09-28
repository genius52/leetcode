package number

import "sort"

func RankTeams(votes []string) string {
	var l int = len(votes)
	var cnt int = len(votes[0])
	var record [26][26]int
	for i := 0;i < l;i++{
		for j := 0;j < cnt;j++{
			record[votes[i][j] - 'A'][j]++
		}
	}
	var data []uint8
	for j := 0;j < cnt;j++{
		data = append(data,votes[0][j])
	}
	sort.Slice(data, func(i, j int) bool {
		for k := 0;k < 26;k++{
			if record[data[i] - 'A'][k] > record[data[j] - 'A'][k]{
				return true
			}else if record[data[i] - 'A'][k] < record[data[j] - 'A'][k]{
				return false
			}
		}
		return data[i] < data[j]
	})
	var res string
	for i := 0;i < len(data);i++{
		res += string(data[i])
	}
	return res
}