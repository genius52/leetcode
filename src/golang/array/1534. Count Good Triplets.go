package array

import "math"

//0 <= i < j < k < arr.length
//|arr[i] - arr[j]| <= a
//|arr[j] - arr[k]| <= b
//|arr[i] - arr[k]| <= c

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var l int = len(arr)
	var res int = 0
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			if(int(math.Abs(float64(arr[i] - arr[j]))) > a){
				continue
			}
			for k := j + 1;k < l;k++{
				if(int(math.Abs(float64(arr[j] - arr[k]))) <= b && int(math.Abs(float64(arr[i] - arr[k]))) <= c){
					res++
				}
			}
		}
	}
	return res
}
