package number

import (
	"math"
	"strconv"
)

//输入：queries = [1,2,3,4,5,90], intLength = 3
//输出：[101,111,121,131,141,999]
func KthPalindrome(queries []int, intLength int) []int64{
	var l int = len(queries)
	var res []int64 = make([]int64,l)
	var data []int = make([]int,intLength)
	data[0] = 1
	data[intLength - 1] = 1
	var base int64 = int64(math.Pow10((intLength - 1)/2))//len = 3,base = 10
	var LIMIT int64 = base * 9//len = 3,max = 90
	for i := 0;i < l;i++{
		if int64(queries[i]) > LIMIT{
			res[i] = -1
			continue
		}
		val := base + int64(queries[i]) - 1//1001的10，12321的123
		var s string = strconv.FormatInt(val,10)
		var idx int = len(s) - 1
		if (intLength | 1) == intLength{
			idx--
		}
		for ;idx >= 0;idx--{
			s += string(s[idx])
		}
		//var n int64 = 0
		n,_ := strconv.ParseInt(s, 10, 64)
		res[i] = n
	}
	return res
}

//func get_next(data *[]int,l int)bool{
//	//var l int = len(*data)
//	var idx int = l / 2
//	for idx >= 0{
//		if (*data)[idx] != 9{
//			break
//		}
//		idx--
//	}
//	if idx < 0{
//		return false
//	}
//	pre1 := (*data)[idx]
//	pre2 := (*data)[l - 1 - idx]
//	(*data)[idx] = pre1 + 1
//	(*data)[l - 1 - idx] = pre2 + 1
//	for i := idx + 1;i < l - 1 - idx;i++{
//		(*data)[i] = 0
//	}
//	return true
//}
//
//func KthPalindrome(queries []int, intLength int) []int64 {
//	var l int = len(queries)
//	var record [][]int = make([][]int,l)
//	for i := 0;i < l;i++{
//		record[i] = []int{queries[i],i}
//	}
//	sort.Slice(record, func(i, j int) bool {
//		return record[i][0] < record[j][0]
//	})
//	var data []int = make([]int,intLength)
//	data[0] = 1
//	data[intLength - 1] = 1
//	var idx int = 1
//	var res []int64 = make([]int64,l)
//	for i := 0;i < l;i++{
//		var overflow bool = false
//		for idx < record[i][0]{
//			ret := get_next(&data,intLength)
//			if !ret{
//				overflow = true
//				break
//			}
//			idx++
//		}
//		if overflow{
//			for j := i;j < l;j++{
//				res[record[j][1]] = -1
//			}
//			break
//		}else{
//			var n int64 = 0
//			for i := 0;i < intLength;i++{
//				n *= 10
//				n += int64(data[i])
//			}
//			res[record[i][1]] = n
//		}
//	}
//	return res
//}