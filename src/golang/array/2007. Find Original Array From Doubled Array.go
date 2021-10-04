package array

import "sort"

func FindOriginalArray(changed []int) []int {
	var l int = len(changed)
	if (l | 1) == l{
		return []int{}
	}
	var record map[int]int = make(map[int]int)
	var keys []int
	for i := 0;i < l;i++{
		record[changed[i]]++
	}
	for k,_ := range record{
		keys = append(keys,k)
	}
	sort.Ints(keys)
	var res []int
	for _,key := range keys{//找当前key 2倍的数字，找不到则
		if cnt1,ok1 := record[key];ok1{
			if key == 0{
				if (cnt1 | 1) == cnt1{
					return []int{}
				}else{
					for i := 0;i < cnt1/2;i++{
						res = append(res,0)
					}
				}
				continue
			}
			dup := key * 2
			if cnt2,ok2 := record[dup];ok2{
				if cnt1 > cnt2{
					return []int{}
				}else if cnt1 < cnt2{
					record[dup] -= cnt1
				}else{
					delete(record,dup)
				}
				for i := 0;i < cnt1;i++{
					res = append(res,key)
				}
			}else{
				return []int{}
			}
		}
	}
	return res
}