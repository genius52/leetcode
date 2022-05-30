package array

import "sort"

func LargestWordCount(messages []string, senders []string) string {
	var l int = len(messages)
	var record map[string]int = make(map[string]int)
	var max_cnt int = 0
	for i := 0;i < l;i++{
		var cnt int = 1
		var idx int = 0
		for idx < len(messages[i]){
			if messages[i][idx] == ' '{
				cnt++
			}
			idx++
		}
		record[senders[i]] += cnt
		if record[senders[i]] > max_cnt{
			max_cnt = record[senders[i]]
		}
	}
	var data []string
	for k,v := range record{
		if v == max_cnt{
			data = append(data,k)
		}
	}
	sort.Strings(data)
	return data[len(data) - 1]
}