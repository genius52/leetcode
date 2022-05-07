package list_queue

import "sort"

type time_val struct {
	timestamp int
	val string
}

type TimeMap struct {
	data map[string][]time_val
}

func Constructor981() TimeMap {
	var obj TimeMap
	obj.data = make(map[string][]time_val)
	return obj
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	this.data[key] = append(this.data[key],time_val{timestamp: timestamp,val: value})
}

//String get(String key, int timestamp)
//返回先前调用set(key, value, timestamp_prev)所存储的值，其中timestamp_prev <= timestamp 。
//如果有多个这样的值，则返回对应最大的timestamp_prev的那个值。
//如果没有值，则返回空字符串（""）。

func (this *TimeMap) Get(key string, timestamp int) string {
	if _,ok := this.data[key];ok{
		var l int = len(this.data[key])
		if this.data[key][0].timestamp > timestamp{
			return ""
		}
		if this.data[key][l - 1].timestamp <= timestamp{
			return this.data[key][l - 1].val
		}
		find_idx := sort.Search(l, func(i int) bool {
			return this.data[key][i].timestamp > timestamp
		})
		if find_idx - 1 >= 0 && find_idx - 1 < l{
			return this.data[key][find_idx - 1].val
		}
	}
	return ""
}