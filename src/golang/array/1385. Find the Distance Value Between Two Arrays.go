package array

import "math"

//Input: arr1 = [4,5,8], arr2 = [10,9,1,8], d = 2
//Output: 2
//Explanation:
//For arr1[0]=4 we have:
//|4-10|=6 > d=2
//|4-9|=5 > d=2
//|4-1|=3 > d=2
//|4-8|=4 > d=2
//For arr1[1]=5 we have:
//|5-10|=5 > d=2
//|5-9|=4 > d=2
//|5-1|=4 > d=2
//|5-8|=3 > d=2
//For arr1[2]=8 we have:
//|8-10|=2 <= d=2
//|8-9|=1 <= d=2
//|8-1|=7 > d=2
//|8-8|=0 <= d=2
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var res int = 0
	for i := 0;i < len(arr1);i++{
		var pass bool = true
		for j:= 0;j < len(arr2);j++{
			if int(math.Abs(float64(arr1[i] - arr2[j]))) <= d{
				pass = false
				break
			}
		}
		if pass{
			res++
		}
	}
	return res
}
