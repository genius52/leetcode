package array

import (
	"math"
	"sort"
)

func squareful(num int)bool{
	n := int(math.Sqrt(float64(num)))
	return (n * n) == num
}

func perm_numSquarefulPerms(nums []int,l int,pos int,res *int){
	if pos >= l{
		*res++
		return
	}
	var visited map[int]bool = make(map[int]bool)
	for i := pos;i < l;i++{
		if _,ok := visited[nums[i]];ok{
			continue
		}
		visited[nums[i]] = true
		nums[i],nums[pos] = nums[pos],nums[i]
		if pos == 0{
			perm_numSquarefulPerms(nums,l,pos + 1,res)
		}else{
			if squareful(nums[pos] + nums[pos - 1]){
				perm_numSquarefulPerms(nums,l,pos + 1,res)
			}
		}
		nums[i],nums[pos] = nums[pos],nums[i]
	}
}

func NumSquarefulPerms(nums []int) int {
	var l int = len(nums)
	sort.Ints(nums)
	var res int = 0
	perm_numSquarefulPerms(nums,l,0,&res)
	return res
}

//func numSquarefulPerms(nums []int) int {
//	var record map[int]int = make(map[int]int)
//	for _,n := range nums{
//		if _,ok := record[n];ok{
//			record[n]++
//		}else{
//			record[n] = 1
//		}
//	}
//	var res int = 0
//	for n,cnt := range record{
//
//	}
//}