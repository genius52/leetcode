package number

import (
	"sort"
	"time"
)

func AlertNames(keyName []string, keyTime []string) []string {
	var l int = len(keyName)
	var record map[string][]string = make(map[string][]string)
	for i := 0;i < l;i++{
		record[keyName[i]] = append(record[keyName[i]],keyTime[i])
	}
	var res []string
	for name,times := range record{
		sort.Strings(times)
		if len(times) <= 2{
			continue
		}
		var cur_len int = len(times)
		var warning bool = false
		for i := 1;i < cur_len - 1;i++{
			prev := times[i - 1]
			next := times[i + 1]
			t1,_ := time.ParseInLocation("15:04", prev,time.Local)
			t2,_ := time.ParseInLocation("15:04",next,time.Local)
			dur := t2.Sub(t1)
			if  dur.Minutes() <= 60{
				warning = true
				break
			}
		}
		if warning{
			res = append(res,name)
		}
	}
	sort.Strings(res)
	return res
}