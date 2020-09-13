package array

import "sort"

//Input: arr = [1,2,3,4,5], k = 4, x = 3
//Output: [1,2,3,4]
func FindClosestElements(arr []int, k int, x int) []int {
	var l int = len(arr)
	if x < arr[0]{
		return arr[0:k]
	}
	if x > arr[l - 1]{
		return arr[l - k:]
	}
	var res []int
	var pos int = 0
	for arr[pos] < x{
		pos++
	}
	head := pos - 1
	tail := pos
	for k > 0{
		if head < 0{
			res = append(res,arr[tail])
			tail++
		}else if tail >= l{
			res = append(res,arr[head])
			head--
		}else{
			if arr[tail] - x < x - arr[head]{
				res = append(res,arr[tail])
				tail++
			}else{
				res = append(res,arr[head])
				head--
			}
		}
		k--
	}
	sort.Ints(res)
	return res
}