package array

import "sort"

func RelativeSortArray2(arr1 []int, arr2 []int) []int{
	var record map[int]int = make(map[int]int)
	var l int = len(arr1)
	for i := 0;i < l;i++{
		if _,ok := record[arr1[i]];!ok{
			record[arr1[i]] = 1
		}else{
			record[arr1[i]]++
		}
	}

	var res []int = make([]int,l)
	var idx int = 0
	for _,n := range arr2{
		cnt := record[n]
		for i := 0;i < cnt;i++{
			res[idx] = n
			idx++
		}
		delete(record,n)
	}
	var keys []int
	for n,_ := range record{
		keys = append(keys,n)
	}
	sort.Ints(keys)
	for i := 0;i < len(keys);i++{
		cnt := record[keys[i]]
		for j := 0;j < cnt;j++{
			res[idx] = keys[i]
			idx++
		}
	}
	return res
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var res []int
	for _,val2 := range arr2{
		for _,val1 := range arr1{
			if val1 == val2{
				res = append(res, val1)
			}
		}
	}
	var rest []int
	for _,val1 := range arr1{
		var find bool = false
		for _,val2 := range arr2{
			if val1 == val2{
				find = true
				break
			}
		}
		if !find {
			rest = append(rest,val1)
		}
	}
	sort.Ints(rest)
	res = append(res,rest...)
	return res
}